// SPDX-License-Identifier: GPL-3.0
pragma solidity 0.8.18;

import "@openzeppelin/contracts/utils/cryptography/ECDSA.sol";
import "@openzeppelin/contracts/utils/cryptography/SignatureChecker.sol";
import "@openzeppelin/contracts/utils/Address.sol";
import "@openzeppelin/contracts/token/ERC721/IERC721.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";

contract Wallet {
    error NotOwnerAuthorized();
    error NotManagerAuthorized();
    error InvalidInput();
    error AlreadyInitialManager();
    error InvalidUserSignature();
    error InvalidEmailSignature();
    error InvalidCodeInput();
    error NotPayeesetted();
    error TimelockInsufficientDelay(uint256 delay, uint256 minDelay);
    error ENotEnoughBalance(uint256 balance);
    error ENotEnoughTokenBalance(uint256 balance);
    error InvalidPayeeTime();

    event emailerror(string indexed inputemail, string indexed storedemail);
    event EthTransPayee(
        address indexed payee,
        address indexed to,
        uint256 indexed amount
    );
    event TokenTransPayee(
        address indexed payee,
        address indexed tokencontract,
        address indexed to,
        uint256 amount
    );
    event NFTTransPayee(
        address indexed payee,
        address indexed tokencontract,
        address indexed to,
        uint256 tokenID
    );

    mapping(string => UserInfo) userinfo;
    bool initialized;
    string email = "";
    mapping(address => PayEthOrder) payEthinfo;
    mapping(address => PayTokenOrder) payTokeninfo;
    mapping(address => PayNFTOrder) payNFTinfo;
    uint256 _minDelay = 300;

    struct UserInfo {
        string email;
        uint256 email_code;
        address owner_address;
        address signaddress;
        address manager;
    }

    struct PayEthOrder {
        uint256 delay;
        uint256 amount;
        address to;
    }
    struct PayTokenOrder {
        address contractaddress;
        uint256 delay;
        uint256 amount; // approve(spender,amount)
        address to;
    }
    struct PayNFTOrder {
        address contractaddress;
        uint256 delay;
        uint256 tokenID; //approve(address to, uint256 tokenId)
        address to;
    }

    modifier onlyOwner() {
        if (msg.sender != owner()) revert NotOwnerAuthorized();
        _;
    }

    modifier onlyEthTransOwn() {
        if (payEthinfo[msg.sender].delay == 0) revert NotPayeesetted();
        _;
    }
    modifier onlyTokenTransOwn() {
        if (payTokeninfo[msg.sender].delay == 0) revert NotPayeesetted();
        _;
    }
    modifier onlyNFTTransOwn() {
        if (payNFTinfo[msg.sender].delay == 0) revert NotPayeesetted();
        _;
    }

    modifier onlyManager() {
        if (msg.sender != userinfo[email].manager)
            revert NotManagerAuthorized();
        _;
    }

    function getMinDelay() public view virtual returns (uint256) {
        return _minDelay;
    }

    function setMinDelay(uint256 newdelay) public virtual onlyManager {
        _minDelay = newdelay;
    }

    function setEthTransPayee(
        uint256 amount,
        address payee,
        address to,
        uint256 _delay,
        string calldata hash,
        bytes calldata signature
    ) public onlyManager {
        if (!isValidUserSignature(hash, signature))
            revert InvalidUserSignature();
        if (address(this).balance < amount) {
            revert ENotEnoughBalance(address(this).balance);
        }
        uint256 minDelay = getMinDelay();
        if (_delay < minDelay) {
            revert TimelockInsufficientDelay(_delay, minDelay);
        }
        PayEthOrder memory payeeorder = PayEthOrder({
            delay: block.timestamp + _delay,
            amount: amount,
            to: to
        });
        payEthinfo[payee] = payeeorder;
        emit EthTransPayee(payee, to, amount);
    }

    function setTokenTransPayee(
        address tokencontract,
        uint256 amount,
        address to,
        address payee,
        uint256 _delay,
        string calldata hash,
        bytes calldata signature
    ) public onlyManager {
        if (!isValidUserSignature(hash, signature))
            revert InvalidUserSignature();
        if (IERC20(tokencontract).balanceOf(address(this)) < amount) {
            revert ENotEnoughTokenBalance(
                IERC20(tokencontract).balanceOf(address(this))
            );
        }
        uint256 minDelay = getMinDelay();
        if (_delay < minDelay) {
            revert TimelockInsufficientDelay(_delay, minDelay);
        }
        PayTokenOrder memory payeeorder = PayTokenOrder({
            contractaddress: tokencontract,
            delay: block.timestamp + _delay,
            amount: amount, // approve(spender,amount)
            to: to
        });
        payTokeninfo[payee] = payeeorder;
        emit TokenTransPayee(payee, tokencontract, to, amount);
    }

    function setNFTTransPayee(
        address tokencontract,
        uint256 tokenID,
        address to,
        address payee,
        uint256 _delay,
        string calldata hash,
        bytes calldata signature
    ) public onlyManager {
        if (!isValidUserSignature(hash, signature))
            revert InvalidUserSignature();
        if (IERC721(tokencontract).ownerOf(tokenID) != address(this)) {
            revert ENotEnoughTokenBalance(
                IERC20(tokencontract).balanceOf(address(this))
            );
        }
        uint256 minDelay = getMinDelay();
        if (_delay < minDelay) {
            revert TimelockInsufficientDelay(_delay, minDelay);
        }
        PayNFTOrder memory payeeorder = PayNFTOrder({
            contractaddress: tokencontract,
            delay: block.timestamp + _delay,
            tokenID: tokenID,
            to: to
            //approve(address to, uint256 tokenId)
        });
        payNFTinfo[payee] = payeeorder;
        emit NFTTransPayee(payee, tokencontract, to, tokenID);
    }

    function resetOrforgetPassword(
        address _newaddress,
        string memory _email,
        uint256 _code,
        bytes calldata emailsignature,
        string calldata hash,
        bytes calldata signature
    ) public onlyManager {
        if (!equal(_email, email)) {
            emit emailerror(_email, email);
        }
        string memory _email_code = concatStrings(_email, _code);
        if (!isValidManagerSignature(_email_code, emailsignature))
            revert InvalidEmailSignature();
        if (!isValidUserSignature(hash, signature))
            revert InvalidUserSignature();
        userinfo[_email].email_code = _code;
        userinfo[_email].owner_address = _newaddress;
    }

    function executeCall(
        address to,
        uint256 value,
        bytes calldata _calldata
    ) external payable onlyOwner returns (bytes memory) {
        return Address.functionCallWithValue(to, _calldata, value);
    }

    function executeEthTrans() external payable onlyEthTransOwn {
        PayEthOrder memory payeeinfo = payEthinfo[msg.sender];
        if (payeeinfo.delay > block.timestamp) revert InvalidPayeeTime();
        Address.sendValue(payable(payeeinfo.to), payeeinfo.amount);
        delete payEthinfo[msg.sender];
    }

    function executeTokenTrans() external payable onlyTokenTransOwn {
        PayTokenOrder memory payeeinfo = payTokeninfo[msg.sender];
        if (payeeinfo.delay > block.timestamp) revert InvalidPayeeTime();
        require(
            IERC20(payeeinfo.contractaddress).transfer(
                payeeinfo.to,
                payeeinfo.amount
            ),
            "Transfer_Faliled"
        );
        delete payTokeninfo[msg.sender];
    }

    function executeNFTTrans() external payable onlyNFTTransOwn {
        PayNFTOrder memory payeeinfo = payNFTinfo[msg.sender];
        if (payeeinfo.delay > block.timestamp) revert InvalidPayeeTime();
        IERC721(payeeinfo.contractaddress).transferFrom(
            address(this),
            payeeinfo.to,
            payeeinfo.tokenID
        );
        delete payNFTinfo[msg.sender];
    }

    function resetManaget(address _manager) public onlyManager {
        if (_manager == address(0)) revert InvalidInput();
        userinfo[email].manager = _manager;
    }

    function resetSignAddress(address _signaddress) external onlyOwner {
        if (_signaddress == address(0)) revert InvalidInput();
        userinfo[email].signaddress = _signaddress;
    }

    function initData(
        address _manager,
        address _signaddress,
        string memory _email,
        uint256 delay
    ) external {
        if (!equal(email, "")) revert AlreadyInitialManager();
        email = _email;

        if (userinfo[_email].manager != address(0))
            revert AlreadyInitialManager();
        if (userinfo[_email].signaddress != address(0)) revert InvalidInput();

        UserInfo memory _userinfo = UserInfo({
            email: _email,
            email_code: 0,
            owner_address: owner(),
            signaddress: _signaddress,
            manager: _manager
        });
        userinfo[_email] = _userinfo;
        initialized = true;
        _minDelay=delay;
    }

    function equal(string memory a, string memory b)
        internal
        pure
        returns (bool)
    {
        return
            bytes(a).length == bytes(b).length &&
            keccak256(bytes(a)) == keccak256(bytes(b));
    }

    function owner() public view returns (address) {
        if (initialized) {
            return userinfo[email].owner_address;
        }
        bytes memory footer = new bytes(0x20);
        assembly {
            extcodecopy(address(), add(footer, 0x20), 0x2d, 0x20)
        }
        return abi.decode(footer, (address));
    }

    function isValidUserSignature(
        string calldata _veridata,
        bytes calldata signature
    ) public view returns (bool) {
        bytes32 _msghash = getMessageHash(_veridata);
        address _owner = userinfo[email].signaddress;
        return isValidSignature(_owner, _msghash, signature);
    }

    function isValidManagerSignature(
        string memory _veridata,
        bytes calldata signature
    ) public view returns (bool) {
        bytes32 _msghash = getMessageHash(_veridata);
        address _manager = userinfo[email].manager;
        return isValidSignature(_manager, _msghash, signature);
    }

    function getMessageHash(string memory str) internal pure returns (bytes32) {
        bytes32 _msgHash = keccak256(abi.encodePacked(str));
        return ECDSA.toEthSignedMessageHash(_msgHash);
    }

    function isValidSignature(
        address _owner,
        bytes32 hash,
        bytes memory signature
    ) internal view returns (bool) {
        // _handleOverrideStatic();

        return SignatureChecker.isValidSignatureNow(_owner, hash, signature);
    }

    function concatStrings(string memory a, uint256 b)
        internal
        pure
        returns (string memory)
    {
        return string(abi.encodePacked(a, b));
    }

    fallback() external payable {}

    receive() external payable {}
}

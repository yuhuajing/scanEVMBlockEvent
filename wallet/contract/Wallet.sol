// SPDX-License-Identifier: GPL-3.0
pragma solidity 0.8.18;

import "@openzeppelin/contracts/utils/cryptography/ECDSA.sol";
import "@openzeppelin/contracts/utils/cryptography/SignatureChecker.sol";
import "@openzeppelin/contracts/utils/Address.sol";
import "@openzeppelin/contracts/token/ERC721/IERC721.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";

contract Wallet {
    error NotPayeeAuthorized();
    error NotOwnerAuthorized();
    error NotManagerAuthorized();
    error InvalidInput();
    error AlreadyInitialManager();
    error AlreadyInitialzed();
    error NotOwnerOnlyEmail();
    error InvalidUserSignature();
    error InvalidEmailSignature();
    error InvalidCodeInput();
    error NotPayeesetted();
    error TimelockInsufficientDelay(uint256 delay, uint256 minDelay);
    error ENotEnoughBalance(uint256 balance);
    error ENotEnoughTokenBalance(uint256 balance);
    error InvalidPayeeTime();
    error AlreadyHasPendingOrder();

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
    string useremail;
    mapping(string => PayEthOrder) payEthinfo;
    mapping(string => PayTokenOrder) payTokeninfo;
    mapping(string => PayNFTOrder) payNFTinfo;
    uint256 _minDelay = 300;

    struct UserInfo {
        string email;
        uint256 email_code;
        address owner;
        address signer;
        address manager;
    }

    struct PayEthOrder {
        uint256 delay;
        address payee;
        address to;
        uint256 amount;
    }
    struct PayTokenOrder {
        address contractaddress;
        uint256 delay;
        address payee;
        uint256 amount;
        address to;
    }
    struct PayNFTOrder {
        address contractaddress;
        uint256 delay;
        address payee;
        uint256 tokenID;
        address to;
    }

    modifier onlyOwner() {
        if (msg.sender != userinfo[email()].owner) revert NotOwnerAuthorized();
        _;
    }

    modifier onlyEthTransOwn() {
        PayEthOrder memory payeeinfo = payEthinfo[useremail];
        if (payeeinfo.delay == 0) revert NotPayeesetted();
        if (payeeinfo.payee != msg.sender) revert NotPayeeAuthorized();
        _;
    }
    modifier onlyTokenTransOwn() {
        PayTokenOrder memory payeeinfo = payTokeninfo[useremail];
        if (payeeinfo.delay == 0) revert NotPayeesetted();
        if (payeeinfo.payee != msg.sender) revert NotPayeeAuthorized();
        _;
    }
    modifier onlyNFTTransOwn() {
        PayNFTOrder memory payeeinfo = payNFTinfo[useremail];
        if (payeeinfo.delay == 0) revert NotPayeesetted();
        if (payeeinfo.payee != msg.sender) revert NotPayeeAuthorized();
        _;
    }

    modifier onlyManager() {
        if (msg.sender != userinfo[email()].manager)
            revert NotManagerAuthorized();
        _;
    }

    function getMinDelay() public view virtual returns (uint256) {
        return _minDelay;
    }

    function setMinDelay(uint256 newdelay) public virtual onlyManager {
        _minDelay = newdelay;
    }

    function getSigner() public view virtual returns (address) {
        return userinfo[email()].signer;
    }

    function getManager() public view virtual returns (address) {
        return userinfo[email()].manager;
    }

    function setEthTransPayee(
        uint256 amount,
        address payee,
        address to,
        uint256 _delay,
        string calldata hash,
        bytes calldata signature
    ) public onlyManager {
        PayEthOrder memory payeeinfo = payEthinfo[useremail];
        if (payeeinfo.delay != 0 && payeeinfo.delay >= block.timestamp)
            revert AlreadyHasPendingOrder();
        if (!isValidUserSignature(hash, signature))
            revert InvalidUserSignature();
        if (address(this).balance < amount) {
            revert ENotEnoughBalance(address(this).balance);
        }
        uint256 minDelay = getMinDelay();
        if (_delay < minDelay) {
            revert TimelockInsufficientDelay(_delay, minDelay);
        }
        if (payeeinfo.delay != 0 && payeeinfo.delay < block.timestamp) {
            payeeinfo.delay = _delay;
            payeeinfo.payee = payee;
            payeeinfo.to = to;
            payeeinfo.amount = amount;
        } else if (payeeinfo.delay == 0) {
            PayEthOrder memory payeeorder = PayEthOrder({
                payee: payee,
                delay: block.timestamp + _delay,
                amount: amount,
                to: to
            });
            payEthinfo[useremail] = payeeorder;
        }
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
        PayTokenOrder memory payeeinfo = payTokeninfo[useremail];
        if (payeeinfo.delay != 0 && payeeinfo.delay >= block.timestamp)
            revert AlreadyHasPendingOrder();
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
        if (payeeinfo.delay != 0 && payeeinfo.delay < block.timestamp) {
            payeeinfo.delay = _delay;
            payeeinfo.payee = payee;
            payeeinfo.to = to;
            payeeinfo.amount = amount;
            payeeinfo.contractaddress = tokencontract;
        } else if (payeeinfo.delay == 0) {
            PayTokenOrder memory payeeorder = PayTokenOrder({
                payee: payee,
                contractaddress: tokencontract,
                delay: block.timestamp + _delay,
                amount: amount, // approve(spender,amount)
                to: to
            });
            payTokeninfo[useremail] = payeeorder;
        }
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
        PayNFTOrder memory payeeinfo = payNFTinfo[useremail];
        if (payeeinfo.delay != 0 && payeeinfo.delay >= block.timestamp)
            revert AlreadyHasPendingOrder();

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
        if (payeeinfo.delay != 0 && payeeinfo.delay < block.timestamp) {
            payeeinfo.delay = _delay;
            payeeinfo.payee = payee;
            payeeinfo.to = to;
            payeeinfo.tokenID = tokenID;
            payeeinfo.contractaddress = tokencontract;
        } else if (payeeinfo.delay == 0) {
            PayNFTOrder memory payeeorder = PayNFTOrder({
                payee: payee,
                contractaddress: tokencontract,
                delay: block.timestamp + _delay,
                tokenID: tokenID,
                to: to
                //approve(address to, uint256 tokenId)
            });
            payNFTinfo[useremail] = payeeorder;
        }
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
        if (!equal(_email, email())) {
            emit emailerror(_email, email());
        }
        string memory _email_code = concatStrings(_email, _code);
        if (!isValidManagerSignature(_email_code, emailsignature))
            revert InvalidEmailSignature();
        if (!isValidUserSignature(hash, signature))
            revert InvalidUserSignature();
        userinfo[_email].email_code = _code;
        userinfo[_email].owner = _newaddress;
    }

    function executeCall(
        address to,
        uint256 value,
        bytes calldata _calldata
    ) external payable onlyOwner returns (bytes memory) {
        return Address.functionCallWithValue(to, _calldata, value);
    }

    function executeEthTrans() external payable onlyEthTransOwn {
        PayEthOrder memory payeeinfo = payEthinfo[useremail];
        if (payeeinfo.delay < block.timestamp) revert InvalidPayeeTime();
        Address.sendValue(payable(payeeinfo.to), payeeinfo.amount);
        //delete payEthinfo[useremail];
    }

    function executeTokenTrans() external payable onlyTokenTransOwn {
        PayTokenOrder memory payeeinfo = payTokeninfo[useremail];
        if (payeeinfo.delay < block.timestamp) revert InvalidPayeeTime();
        require(
            IERC20(payeeinfo.contractaddress).transfer(
                payeeinfo.to,
                payeeinfo.amount
            ),
            "Transfer_Faliled"
        );
        // delete payEthinfo[useremail];
    }

    function executeNFTTrans() external payable onlyNFTTransOwn {
        PayNFTOrder memory payeeinfo = payNFTinfo[useremail];
        if (payeeinfo.delay < block.timestamp) revert InvalidPayeeTime();
        IERC721(payeeinfo.contractaddress).transferFrom(
            address(this),
            payeeinfo.to,
            payeeinfo.tokenID
        );
        //delete payEthinfo[useremail];
    }

    function resetManaget(address _manager) public onlyManager {
        if (_manager == address(0)) revert InvalidInput();
        userinfo[email()].manager = _manager;
    }

    function resetSignAddress(address _signaddress) external onlyOwner {
        if (_signaddress == address(0)) revert InvalidInput();
        userinfo[email()].signer = _signaddress;
    }

    function initData(
        address _owneraddress,
        address _manager,
        address _signaddress,
        uint256 delay
    ) external {
        if (initialized) revert AlreadyInitialzed();
        useremail = email();
        UserInfo memory _userinfo = UserInfo({
            email: useremail,
            email_code: 0,
            owner: _owneraddress,
            signer: _signaddress,
            manager: _manager
        });
        userinfo[useremail] = _userinfo;
        _minDelay = delay;
        initialized = true;
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

    function email() public view returns (string memory) {
        if (initialized) {
            return useremail;
        }
        bytes memory footer = new bytes(0x20);
        assembly {
            extcodecopy(address(), add(footer, 0x20), 0x2d, 0x20)
        }
        return convertByte32ToString(abi.decode(footer, (bytes32)));
    }

    function convertByte32ToString(bytes32 _bytes32)
        internal
        pure
        returns (string memory)
    {
        bytes memory bytesArray = new bytes(32);
        for (uint256 i; i < 32; i++) {
            bytesArray[i] = _bytes32[i];
        }
        return string(bytesArray);
    }

    function isValidUserSignature(
        string calldata _veridata,
        bytes calldata signature
    ) public view returns (bool) {
        bytes32 _msghash = getMessageHash(_veridata);
        address _owner = userinfo[email()].signer;
        return isValidSignature(_owner, _msghash, signature);
    }

    function isValidManagerSignature(
        string memory _veridata,
        bytes calldata signature
    ) public view returns (bool) {
        bytes32 _msghash = getMessageHash(_veridata);
        address _manager = userinfo[email()].manager;
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

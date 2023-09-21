// SPDX-License-Identifier: GPL-3.0

pragma solidity 0.8.18;

import "@openzeppelin/contracts/utils/cryptography/SignatureChecker.sol";

contract Wallet {
    error NotOwnerAuthorized();
    error NotManagerAuthorized();
    error InvalidInput();
    error AlreadyInitialManager();
    error InvalidQAInput();
    error InvalidUserSignature();
    error InvalidCodeInput();
    event emailerror(string indexed inputemail, string indexed storedemail);
   // event QAerror(string indexed inputQA, string indexed storedQA);
    
    mapping(string => UserInfo) userinfo;
    mapping(string => mapping(uint256 => bool)) email_code;
   // bool resetOwner;
    bool initialized;
    bool resetQA;
    string QA;
    string email = "";

    struct UserInfo {
        string email;
        uint256 email_code;
        address owner_address;
        address signaddress;
        address manager;
    }

    modifier onlyOwner() {
        if (msg.sender != owner()) revert NotOwnerAuthorized();
        _;
    }
    modifier onlyManager() {
        if (msg.sender != userinfo[email].manager)
            revert NotManagerAuthorized();
        _;
    }

    function initData(
        address _manager,
        address _signaddress,
        string memory _email
    ) external {
        if (!equal(email, "")) revert AlreadyInitialManager();
        email = _email;

        if (userinfo[_email].manager != address(0))
            revert AlreadyInitialManager();
        // userinfo[email].manager = _manager;
        if (userinfo[_email].signaddress != address(0)) revert InvalidInput();
        // signaddress = _signaddress;

        UserInfo memory _userinfo = UserInfo({
            email: _email,
            email_code: 0,
            owner_address: owner(),
            signaddress: _signaddress,
            manager: _manager
        });
        userinfo[_email] = _userinfo;
        initialized = true;
    }

    // function initManager(address _manager) external {
    //     if (manager != address(0)) revert AlreadyInitialManager();
    //     manager = _manager;
    // }

    // function initSignAddress(address _signaddress) external {
    //     if (signaddress != address(0)) revert InvalidInput();
    //     signaddress = _signaddress;
    // }

    function resetManaget(address _manager) public onlyManager {
        if (_manager == address(0)) revert InvalidInput();
        //manager = _manager;
        userinfo[email].manager = _manager;
    }

    function resetSignAddress(address _signaddress) external onlyOwner {
        if (_signaddress == address(0)) revert InvalidInput();
        //signaddress = _signaddress;
        userinfo[email].signaddress = _signaddress;
    }

    // function concatStrings(string memory a, string memory b)
    //     internal
    //     pure
    //     returns (string memory)
    // {
    //     return string(abi.encodePacked(a, b));
    // }

    function resetOrforgetPassword(
        //address _address,
        address _newaddress,
        string memory _email,
        uint256 _code,
        bytes32 hash,
        bytes memory signature
    ) public {

        if (!equal(_email, userinfo[_email].email)) {
            emit emailerror(_email, userinfo[_email].email);
        }
        if (!isverified(_email, _code)) revert InvalidCodeInput();
        if (!isValidSignature(hash, signature)) revert InvalidUserSignature(); // revert InvalidQAInput();

        userinfo[_email].email_code = _code;
        userinfo[_email].owner_address = _newaddress;
 
        delete email_code[_email][_code];
       // setowner(_newaddress);
    }

    function executeCall(
        address to,
        uint256 value,
        bytes calldata _calldata
    ) external payable onlyOwner returns (bytes memory) {
        //_incrementNonce();
        return _call(to, value, _calldata);
    }

    function _call(
        address to,
        uint256 value,
        bytes calldata _calldata
    ) internal returns (bytes memory result) {
        bool success;
        (success, result) = to.call{value: value}(_calldata);

        if (!success) {
            assembly {
                revert(add(result, 32), mload(result))
            }
        }
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

    function storeuserinfo(
        address _address,
        string memory _email,
        address _signedaddress,
        address _manager
    ) private {
        UserInfo memory _userinfo = UserInfo({
            email: _email,
            email_code: 0,
            owner_address: _address,
            signaddress: _signedaddress,
            manager: _manager
        });
        userinfo[_email] = _userinfo;
        initialized = true;
    }

    function isverified(string memory _email, uint256 _code)
        internal
        view
        returns (bool)
    {
        return email_code[_email][_code];
    }

    function verifycode(string memory _email, uint256 _code)
        public
        onlyManager
    {
        //Oracle
        email_code[_email][_code] = true;
    }

    function convertByte32ToString(bytes32 _bytes32)
        public
        pure
        returns (string memory)
    {
        bytes memory bytesArray = new bytes(32);
        for (uint256 i; i < 32; i++) {
            bytesArray[i] = _bytes32[i];
        }
        return string(bytesArray);
    }

    function data()
        internal
        view
        returns (
            address _address,
            string memory _email,
            string memory _mixed_password,
            string memory _mixed_question_mixed_answer
        )
    {
        bytes memory footer = new bytes(0x80);
        assembly {
            extcodecopy(address(), add(footer, 0x20), 0x2d, 0x80)
        }
        bytes32 dataemail;
        bytes32 mixed_password;
        bytes32 mixed_question_mixed_answer;
        (_address, dataemail, mixed_password, mixed_question_mixed_answer) = abi
            .decode(footer, (address, bytes32, bytes32, bytes32));
        _email = convertByte32ToString(dataemail);
        _mixed_password = convertByte32ToString(mixed_password);
        _mixed_question_mixed_answer = convertByte32ToString(
            mixed_question_mixed_answer
        );
    }

    // function getQA() public view returns (string memory) {
    //     if (resetQA) {
    //         return QA;
    //     }
    //     (
    //         address addr,
    //         string memory email,
    //         string memory mixed_password,
    //         string memory mixed_question_mixed_answer
    //     ) = data();
    //     return mixed_question_mixed_answer;
    // }

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

    function _isValidSigner(address signer) internal view returns (bool) {
        return signer == owner();
    }

    function isValidSignature(bytes32 hash, bytes memory signature)
        public
        view
        returns (bool)
    {
        bool isValid = SignatureChecker.isValidSignatureNow(
            userinfo[email].signaddress,
            hash,
            signature
        );

        return isValid;
    }

    fallback() external payable {}

    receive() external payable {}
}

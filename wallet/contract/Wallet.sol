// SPDX-License-Identifier: GPL-3.0

pragma solidity 0.8.18;

contract Wallet {
    error NotAuthorized();
    error InvalidInput();

    mapping(address => UserInfo) userinfo;
    mapping(string => mapping(uint256 => bool)) email_code;
    address _owner;
    address manager;
    bool resetOwner;
    bool initialized;

    struct UserInfo {
        string email;
        uint256 email_code;
        address owner_address;
        string node_mixed_password;
        string node_mixed_question;
        string node_mixed_answer;
    }

    modifier onlyOwner() {
        if (msg.sender != owner()) revert NotAuthorized();
        _;
    }
    modifier onlyManager() {
        if (msg.sender != manager) revert NotAuthorized();
        _;
    }

    constructor(address _manager) {
        if (_manager == address(0)) revert InvalidInput();
        manager = _manager;
    }

    function gettestnum()public pure returns(uint256){
        return 5;
    }

    function resetOrforgetPassword(
        address _address,
        address _newaddress,
        string memory _email,
        uint256 _code,
        string memory _mixed_question,
        string memory _mixed_answer
    ) public {
        if (!initialized) {
            (
                address addr,
                string memory email,
                string memory mixed_password,
                string memory mixed_question,
                string memory mixed_answer
            ) = data();
            storeuserinfo(
                addr,
                email,
                mixed_password,
                mixed_question,
                mixed_answer
            );
        }
        UserInfo memory _userinfo = userinfo[_address];
        require(
            isverified(_email, _code) &&
                _address == owner() &&
                equal(_email, _userinfo.email) &&
                equal(_mixed_question, _userinfo.node_mixed_question) &&
                equal(_mixed_answer, _userinfo.node_mixed_answer)
        );
        _userinfo.email_code = _code;
        _userinfo.owner_address = _newaddress;
        // UserInfo memory newuserinfo =_userinfo;
        userinfo[_newaddress] = _userinfo;
        delete userinfo[_address];
        setowner(_newaddress);
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
        string memory _mixed_password,
        string memory _mixed_question,
        string memory _mixed_answer
    ) private {
        UserInfo memory _userinfo = UserInfo({
            email: _email,
            email_code: 0,
            owner_address: _address,
            node_mixed_password: _mixed_password,
            node_mixed_question: _mixed_question,
            node_mixed_answer: _mixed_answer
        });
        userinfo[_address] = _userinfo;
        initialized = true;
    }

    function isverified(string memory _email, uint256 _code)
        public
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

    function data()
        public
        view
        returns (
            address,
            string memory,
            string memory,
            string memory,
            string memory
        )
    {
        bytes memory footer = new bytes(0x60);
        assembly {
            extcodecopy(address(), add(footer, 0x20), 0x4d, 0x60)
        }
        return abi.decode(footer, (address, string, string, string, string));
    }

    function owner() public view returns (address) {
        if (resetOwner) {
            return _owner;
        }
        (
            address _address,
            string memory _email,
            string memory _mixed_password,
            string memory _mixed_question,
            string memory _mixed_answer
        ) = data();
        return _address;
    }

    function setowner(address _addr) private {
        _owner = _addr;
        resetOwner = true;
    }

    function _isValidSigner(address signer) internal view returns (bool) {
        return signer == owner();
    }
}

// SPDX-License-Identifier: GPL-3.0

pragma solidity 0.8.18;

contract Wallet {
    mapping(address => UserInfo) userinfo;
    address public owner;
    address public implementation;

    struct UserInfo {
        string email;
        uint256 email_code;
        address owner_address;
        UserTree usertree;
    }

    struct UserTree {
        string tree_root;
        string tree_node_password;
        string tree_node_question;
        string tree_node_answer;
    }

    function initialize(
        address _address,
        string memory email,
        string memory _mixed_password,
        string memory _mixed_question,
        string memory _mixed_answer,
        address _implementation
    ) external {
        storeuserinfo(
            _address,
            email,
            _mixed_password,
            _mixed_question,
            _mixed_answer
        );
        _setowner(_address);
        _setimplementation(_implementation);
    }

    function merkleroot(
        string memory _email,
        string memory _mixed_password,
        string memory _mixed_question,
        string memory _mixed_answer
    ) public pure returns (string memory) {
        bytes32[] memory bytesArray = new bytes32[](4);
        bytesArray[0] = convertStringToByte32(_email);
        bytesArray[1] = convertStringToByte32(_mixed_password);
        bytesArray[2] = convertStringToByte32(_mixed_question);
        bytesArray[3] = convertStringToByte32(_mixed_answer);
        string memory root = convertByte32ToString(
            computeMerkleRoot(bytesArray)
        );
        return root;
    }

    function storeuserinfo(
        address _address,
        string memory _email,
        string memory _mixed_password,
        string memory _mixed_question,
        string memory _mixed_answer
    ) private {
        string memory root = merkleroot(
            _email,
            _mixed_password,
            _mixed_question,
            _mixed_answer
        );
        UserTree memory _usertree = UserTree({
            tree_root: root,
            tree_node_password: _mixed_password,
            tree_node_question: _mixed_question,
            tree_node_answer: _mixed_answer
        });
        UserInfo memory _userinfo = UserInfo({
            email: _email,
            email_code: 0,
            owner_address: _address,
            usertree: _usertree
        });
        userinfo[_address] = _userinfo;
    }

    function resetOrforgetPassword(
        address _address,
        address _newaddress,
        string memory _email,
        uint256 _code,
        string memory _mixed_question,
        string memory _mixed_answer
    ) public {
        UserInfo memory _userinfo = userinfo[_address];
        string memory _mixed_password = _userinfo.usertree.tree_node_password;
        string memory newroot = merkleroot(
            _email,
            _mixed_password,
            _mixed_question,
            _mixed_answer
        );
        require(equal(newroot, _userinfo.usertree.tree_root));
        _userinfo.email_code = _code;
        _setowner(_newaddress);
    }

    // function getroot(address _address) public view returns (string memory) {
    //     UserInfo memory _userinfo = userinfo[_address];
    //     return _userinfo.usertree.tree_root;
    // }

    function _setowner(address _addr) public {
        owner = _addr;
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

    uint256 private constant size = 32;

    function convertStringToByte32(string memory _texte)
        public
        pure
        returns (bytes32 result)
    {
        assembly {
            result := mload(add(_texte, 32))
        }
    }

    function convertByte32ToString(bytes32 _bytes32)
        public
        pure
        returns (string memory)
    {
        bytes memory bytesArray = new bytes(size);
        for (uint256 i; i < size; i++) {
            bytesArray[i] = _bytes32[i];
        }
        return string(bytesArray);
    }

    function processProof(bytes32[] memory proof, bytes32 leaf)
        internal
        pure
        returns (bytes32)
    {
        bytes32 computedHash = leaf;
        for (uint256 i = 0; i < proof.length; i++) {
            bytes32 proofElement = proof[i];
            if (computedHash <= proofElement) {
                // Hash(current computed hash + current element of the proof)
                computedHash = _efficientHash(computedHash, proofElement);
            } else {
                // Hash(current element of the proof + current computed hash)
                computedHash = _efficientHash(proofElement, computedHash);
            }
        }
        return computedHash;
    }

    function computeMerkleRoot(bytes32[] memory leaves)
        public
        pure
        returns (bytes32)
    {
        require(leaves.length > 0, "No leaves provided");

        if (leaves.length == 1) {
            return leaves[0];
        } else {
            uint256 n = leaves.length;
            bytes32[] memory parents = new bytes32[](n / 2 + (n % 2));

            for (uint256 i = 0; i < n / 2; i++) {
                parents[i] = _efficientHash(leaves[2 * i], leaves[2 * i + 1]);
            }

            if (n % 2 != 0) {
                parents[n / 2] = _efficientHash(leaves[n - 1], leaves[n - 1]);
            }

            return computeMerkleRoot(parents);
        }
    }

    function _efficientHash(bytes32 a, bytes32 b)
        private
        pure
        returns (bytes32 value)
    {
        assembly {
            mstore(0x00, a)
            mstore(0x20, b)
            value := keccak256(0x00, 0x40)
        }
    }

    function _fallback() internal virtual {
        _delegate(implementation);
    }

    function _setimplementation(address _implementation) internal view virtual {
        _implementation = implementation;
    }

    /**
     * @dev Fallback function that delegates calls to the address returned by `_implementation()`. Will run if no other
     * function in the contract matches the call data.
     */
    fallback() external payable virtual {
        _fallback();
    }

    receive() external payable virtual {}

    function _delegate(address _implementation) internal virtual {
        assembly {
            // Copy msg.data. We take full control of memory in this inline assembly
            // block because it will not return to Solidity code. We overwrite the
            // Solidity scratch pad at memory position 0.
            calldatacopy(0, 0, calldatasize())

            // Call the implementation.
            // out and outsize are 0 because we don't know the size yet.
            let result := delegatecall(
                gas(),
                _implementation,
                0,
                calldatasize(),
                0,
                0
            )
            // Copy the returned data.
            returndatacopy(0, 0, returndatasize())

            switch result
            // delegatecall returns 0 on error.
            case 0 {
                revert(0, returndatasize())
            }
            default {
                return(0, returndatasize())
            }
        }
    }
}

contract WalletFactory {
    // mapping(address => mapping(address => address)) public getpair;
    address[] userwallet;

    function addElement(address element) private {
        userwallet.push(element);
    }

    function getElement(uint256 index) public view returns (address) {
        require(index < userwallet.length, "Invalid index");
        return userwallet[index];
    }

    function getLength() public view returns (uint256) {
        return userwallet.length;
    }

    //constructor() payable {}
    function createAccount(
        address _address,
        string memory email,
        string memory _mixed_password,
        string memory _mixed_question,
        string memory _mixed_answer,
        address _implementation,
        uint256 _salt
    ) public returns (address pairAddr) {
        bytes32 salt = keccak256(abi.encodePacked(_salt));
        Wallet wallet = new Wallet{salt: salt}();
        wallet.initialize(
            _address,
            email,
            _mixed_password,
            _mixed_question,
            _mixed_answer,
            _implementation
        );
        pairAddr = address(wallet);
        addElement(pairAddr);
    }

    function calculateAddr(
        uint256 _salt
    ) public view returns (address newAddr) {
        bytes32 salt = keccak256(abi.encodePacked(_salt));
        newAddr = address(
            uint160(
                uint256(
                    keccak256(
                        abi.encodePacked(
                            bytes1(0xff),
                            address(this),
                            salt,
                            keccak256(type(Wallet).creationCode)
                        )
                    )
                )
            )
        );
    }
}

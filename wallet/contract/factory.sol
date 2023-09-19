// SPDX-License-Identifier: MIT
pragma solidity 0.8.18;
import "@openzeppelin/contracts/utils/Create2.sol";

contract Factory {
    mapping(address => mapping(address => address)) public getpair;

    constructor() payable {}

    function getCreationCode(
        address _address,
        string memory email,
        string memory _mixed_password,
        string memory _mixed_question,
        string memory _mixed_answer,
        address _implementation
    )
        internal
        pure
        returns (
            //uint256 _salt
            bytes memory
        )
    {
        return
            abi.encodePacked(
                hex"3d60ad80600a3d3981f3363d3d373d3d3d363d73",
                _implementation,
                hex"5af43d82803e903d91602b57fd5bf3",
                abi.encode(
                    _address,
                    email,
                    _mixed_password,
                    _mixed_question,
                    _mixed_answer
                    //_salt
                )
            );
    }

    function createAccount(
        address _address,
        string memory email,
        string memory _mixed_password,
        string memory _mixed_question,
        string memory _mixed_answer,
        address _implementation,
        uint256 _salt
    ) external returns (address) {
        bytes memory code = getCreationCode(
            _address,
            email,
            _mixed_password,
            _mixed_question,
            _mixed_answer,
            _implementation
        );

        address _account = Create2.computeAddress(
            bytes32(_salt),
            keccak256(code)
        );

        // if (_account.code.length != 0) return _account;

        //  emit AccountCreated(_account, implementation, chainId, tokenContract, tokenId, salt);

        assembly {
            _account := create2(0, add(code, 0x20), mload(code), _salt)
        }
        //if (_account == address(0)) revert AccountCreationFailed();

        return _account;
    }

    function account(
        address _address,
        string memory email,
        string memory _mixed_password,
        string memory _mixed_question,
        string memory _mixed_answer,
        address _implementation,
        uint256 _salt
    ) external view returns (address) {
        bytes32 bytecodeHash = keccak256(
            getCreationCode(
                _address,
                email,
                _mixed_password,
                _mixed_question,
                _mixed_answer,
                _implementation
            )
        );

        return Create2.computeAddress(bytes32(_salt), bytecodeHash);
    }
}

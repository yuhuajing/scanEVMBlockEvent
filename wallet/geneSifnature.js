import { ethers } from "ethers";

const main = async () => {
    const provider = new ethers.JsonRpcProvider(`https://cloudflare-eth.com`);
    const signer = new ethers.Wallet("f07a77cb019764a524dce24cb47ac62bb231b4f0d7bab5f864f603f8cb0e344c", provider)
    console.log(`私钥钱包地址:${signer.address}`)

    // // 创建消息
    // const account = "0x5B38Da6a701c568545dCfcB03FcB875f56beddC4"
    // const tokenId = "0"
    // const msgHash = ethers.solidityPackedKeccak256(
    //     ['address', 'uint256'],
    //     [account, tokenId])
        // // 创建消息
    const str = "helloworld"
    // 等效于Solidity中的keccak256(abi.encodePacked(account, tokenId))
    const msgHash = ethers.solidityPackedKeccak256(
        ['string'],
        [str])
    const messageHashBytes = ethers.getBytes(msgHash)
    const signature = await signer.signMessage(messageHashBytes);
    console.log(`签名：${signature}`) //https://github.com/yuhuajing/solidityLearn/blob/main/smartContract/ECDSA/ECDSA.md
}
main()
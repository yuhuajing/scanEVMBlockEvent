import { ethers } from "ethers";

const main = async () => {
    const provider = new ethers.JsonRpcProvider(`https://cloudflare-eth.com`);
    //Signer 0xd514Ca657E536bB30962A31b22B4F39183328E0F
    signer_prikey = "7174cd9f4f8cd8bd8c91898744ba231b5db50d5191f14819408d51ccf8c6a8c9"
    //Manager 0xf5fBB766074124A574fc9aFaF9c9f139e7efB981
    manager_prikey = "f07a77cb019764a524dce24cb47ac62bb231b4f0d7bab5f864f603f8cb0e344c"
    //Owner 0x156b6c24e78fede687950ba52a0b6b15a2c0ae11
    owner_prikey = "867601ac4dc7028894d3ec525199a5289eeaa9ae38deba3a02511b31ce274901"
    //New_owner 0x9ab95fbf671a3b40f977eb116f948f69b26e663d
    new_owner_prikey = "5567ceafd8404b4d3578f454cd7b78e82c1bdd7711cabdc0b10da72c4d0a24f8"

    const signer = new ethers.Wallet(signer_prikey, provider)
    console.log(`Signer钱包地址:${signer.address}`)
    const signerstr = "signerMessage"
    // 等效于Solidity中的keccak256(abi.encodePacked(account, tokenId))
    const signermsgHash = ethers.solidityPackedKeccak256(
        ['string'],
        [signerstr])
    const signermessageHashBytes = ethers.getBytes(signermsgHash)
    const signersignature = await signer.signMessage(signermessageHashBytes);
    console.log(`Signer签名:${signersignature}`) //https://github.com/yuhuajing/solidityLearn/blob/main/smartContract/ECDSA/ECDSA.md


    const manager = new ethers.Wallet(manager_prikey, provider)
    console.log(`manager钱包地址:${manager.address}`)
    const managerstr = "ManagerMessage"
    // 等效于Solidity中的keccak256(abi.encodePacked(account, tokenId))
    const managermsgHash = ethers.solidityPackedKeccak256(
        ['string'],
        [managerstr])
    const managermessageHashBytes = ethers.getBytes(managermsgHash)
    const managersignature = await signer.signMessage(managermessageHashBytes);
    console.log(`Manager签名:${managersignature}`) //https://github.com/yuhuajing/solidityLearn/blob/main/smartContract/ECDSA/ECDSA.md
}
main()
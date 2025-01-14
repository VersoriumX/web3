require('dotenv').config();
const { ethers } = require('hardhat');

async function main() {
    const poolAddress = 'YOUR_POOL_ADDRESS'; // Replace with your deployed pool address
    const LiquidityPool = await ethers.getContractFactory('LiquidityPool');
    const pool = await LiquidityPool.attach(poolAddress);

    // Example of adding liquidity
    const tokenA = '0x8487B97c91ecC1a03b4907B64Bdeab306B888c0E'; // EthereumX
    const tokenB = '0x7ceB23fD6bC0adD59E62ac25578270cFf1b9f619'; // Wrapped Ethereum
    const amountADesired = ethers.utils.parseEther('1'); // Amount of token A
    const amountBDesired = ethers.utils.parseEther('1'); // Amount of token B

    const tx = await pool.createPool(tokenA, tokenB, amountADesired, amountBDesired);
    await tx.wait();
    console.log('Liquidity added successfully');
}

main()
    .then(() => process.exit(0))
    .catch((error) => {
        console.error(error);
       

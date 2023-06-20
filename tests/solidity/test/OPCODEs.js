require('dotenv').config()
const { expect } = require('chai')
const { sendShieldedTransaction, getProvider } = require("./testUtils")

describe('OPCODE test', () => {
    let contractInstance
    const provider = getProvider()
    const signerPrivateKey = process.env.FIRST_PRIVATE_KEY

    beforeEach(async () => {
        const OpcodesContract = await ethers.getContractFactory('OpCodes')
        contractInstance = await OpcodesContract.deploy({gasLimit: 1000000})
        await contractInstance.deployed()
    })

    it('Should throw invalid op code', async () => {
        let failed = false
        try {
            const tx = await sendShieldedTransaction(
                provider,
                signerPrivateKey,
                contractInstance.address,
                contractInstance.interface.encodeFunctionData("test_invalid", [])
            )
            await tx.wait()
        } catch {
            failed = true
        }

        expect(failed).to.be.true
    })

    it('Should revert', async () => {
        let failed = false
        try {
            const tx = await sendShieldedTransaction(
                provider,
                signerPrivateKey,
                contractInstance.address,
                contractInstance.interface.encodeFunctionData("test_revert", [])
            )
            await tx.wait()
        } catch {
            failed = true
        }

        expect(failed).to.be.true
    })
});
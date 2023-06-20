const {anyValue} = require("@nomicfoundation/hardhat-chai-matchers/withArgs");
const {expect} = require("chai");

describe('Counter', () => {
    let counterContract

    before(async () => {
        const Counter = await ethers.getContractFactory('Counter')
        counterContract = await Counter.deploy()
        await counterContract.deployed()
    })

    it('Should add', async () => {
        const countBefore = await counterContract.counter()

        await expect(counterContract.add())
            .to.emit(counterContract, 'Added')
            .to.emit(counterContract, 'Changed')

        const countAfter = await counterContract.counter()
        expect(countAfter).to.be.equal(countBefore + 1)
    })

    it('Should subtract', async () => {
        const countBefore = await counterContract.counter()

        await expect(counterContract.subtract())
            .to.emit(counterContract, 'Changed')

        const countAfter = await counterContract.counter()
        expect(countAfter).to.be.equal(countBefore - 1)
    })

    it('Should revert correctly', async () => {
        const tx = await counterContract.subtract()
        await expect(tx.wait()).to.be.rejected
    })
})
import FungibleToken from 0x%s
import EngramToken from 0x%s

transaction(to: Address, amount: UFix64) {

    let sentVault: @FungibleToken.Vault

    prepare(signer: AuthAccount) {

       
        let vaultRef = signer.borrow<&EngramToken.Vault>(from: EngramToken.VaultStoragePath)
			?? panic("Could not borrow reference to the owner's Vault!")

    
        self.sentVault <- vaultRef.withdraw(amount: amount)
    }

    execute {

        let recipient = getAccount(to)

        let receiverRef = recipient.getCapability(EngramToken.ReceiverPublicPath)!.borrow<&{FungibleToken.Receiver}>()
			?? panic("Could not borrow receiver reference to the recipient's Vault")

        receiverRef.deposit(from: <-self.sentVault)
    }
}
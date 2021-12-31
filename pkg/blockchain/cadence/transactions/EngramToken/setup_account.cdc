import FungibleToken from 0x%s
import EngramToken from 0x%s

transaction {

    prepare(signer: AuthAccount) {

        if signer.borrow<&EngramToken.Vault>(from: EngramToken.VaultStoragePath) == nil {
            signer.save(<-EngramToken.createEmptyVault(), to: EngramToken.VaultStoragePath)

            signer.link<&EngramToken.Vault{FungibleToken.Receiver}>(
                EngramToken.ReceiverPublicPath,
                target: EngramToken.VaultStoragePath
            )

            signer.link<&EngramToken.Vault{FungibleToken.Balance}>(
                EngramToken.BalancePublicPath,
                target: EngramToken.VaultStoragePath
            )
        }
    }
}
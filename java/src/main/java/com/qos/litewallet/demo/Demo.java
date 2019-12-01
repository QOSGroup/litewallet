package com.qos.litewallet.demo;

import com.alibaba.fastjson.JSON;
import com.alibaba.fastjson.JSONObject;
import com.qos.litewallet.demo.domain.Account;
import com.qos.litewallet.demo.domain.AccountWithMnemonic;
import com.qos.litewallet.demo.response.*;

import java.io.File;
import java.util.Base64;

/**
 * @author wangzhiyong
 * @date 19-11-27下午2:04
 */

public class Demo {

    public static final void main(String[] args) {

        String storagePath = "/tmp/litewallet";
        String walletName = "qos";

        LiteWalletWrapper liteWallet = new LiteWalletWrapper(walletName, storagePath);


        //1. 生成助记词
        ProduceMnemonicResponse mnemonicResponse = liteWallet.ProduceMnemonic();
        System.out.println(mnemonicResponse);

        //2. 创建账户
        String password = "12345678";
        String wrongPassword = "1234567";

        ///使用默认生成的名称及随机的助记词生成账户
        CreateAccountResponse createAccountResponse = liteWallet.CreateAccount(password);
        AccountWithMnemonic accountWithMnemonic = createAccountResponse.getData();

        String address = accountWithMnemonic.getAddress();
        String mnemonic = accountWithMnemonic.getMnemonic();

        //3. 查询账户

        ///使用地址查询账户
        GetAccountResponse getAccountResponse = liteWallet.GetAccount(address);
        System.out.println(getAccountResponse.getData().getAddress());

        ///使用名称查询账户
        getAccountResponse = liteWallet.GetAccountByName("NotExists");
        System.out.println(getAccountResponse.getCode());

        //4. 导出账户
        ExportAccountResponse exportAccountResponse = liteWallet.ExportAccount(address, password);
        String privateKey = exportAccountResponse.getData().getPrivateKey();
        System.out.println(privateKey);

        exportAccountResponse = liteWallet.ExportAccount(address, wrongPassword);
        System.out.println(exportAccountResponse.getMessage());


        //5. 删除用户
        DeleteAccountResponse deleteAccountResponse = liteWallet.DeleteAccount(address, wrongPassword);
        System.out.println(deleteAccountResponse.getMessage());

        //6. 获取所有账户列表
        ListAccountsResponse listAccountsResponse = liteWallet.ListAllAccounts();
        System.out.println(listAccountsResponse);

        ///删除所有账户
        for(Account account : listAccountsResponse.getData()){
            LiteWallet.INSTANCE.DeleteAccount(account.getAddress(), password);
        }


        //7. 导入账户
        ImportAccountResponse importAccountResponse = JSON.parseObject(LiteWallet.INSTANCE.ImportMnemonic(mnemonic, password), ImportAccountResponse.class);
        System.out.println(importAccountResponse);

        //8. 签名

        String message = "你好, QOS钱包";
        String base64Message = new String(Base64.getEncoder().encode(message.getBytes()));

        SignDataResponse signDataResponse1 = liteWallet.Sign(address, password, message);
        SignDataResponse signDataResponse2 = liteWallet.SignBase64(address, password, base64Message);

        System.out.println(signDataResponse1);
        System.out.println(signDataResponse2);

    }


}

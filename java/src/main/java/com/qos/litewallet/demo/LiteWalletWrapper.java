package com.qos.litewallet.demo;

import com.alibaba.fastjson.JSON;
import com.qos.litewallet.demo.response.*;

import java.io.File;

/**
 * @author wangzhiyong
 * @date 19-11-27下午4:27
 */

public class LiteWalletWrapper {

    private String walletName;
    private String walletStoragePath;

    public LiteWalletWrapper(String walletName, String walletStoragePath) {
        this.walletName = walletName;
        this.walletStoragePath = walletStoragePath;
        File walletStoragePathFile = new File(walletStoragePath);
        if (!walletStoragePathFile.exists()) {
            walletStoragePathFile.mkdirs();
        }
        LiteWallet.INSTANCE.InitWallet(walletName, walletStoragePath);
    }


    //生成助记词
    public ProduceMnemonicResponse ProduceMnemonic() {
        return JSON.parseObject(LiteWallet.INSTANCE.ProduceMnemonic(), ProduceMnemonicResponse.class);
    }

    //创建账户
    public CreateAccountResponse CreateAccount(String password) {
        return JSON.parseObject(LiteWallet.INSTANCE.CreateAccount(password), CreateAccountResponse.class);
    }

    //使用指定名称创建账户
    public CreateAccountResponse CreateAccountWithName(String name, String password) {
        return JSON.parseObject(LiteWallet.INSTANCE.CreateAccountWithName(name, password), CreateAccountResponse.class);
    }

    //使用指定名称和助记词创建账户
    public CreateAccountResponse CreateAccountWithMnemonic(String name, String password, String mnemonic) {
        return JSON.parseObject(LiteWallet.INSTANCE.CreateAccountWithMnemonic(name, password, mnemonic), CreateAccountResponse.class);
    }

    //查询账户
    public GetAccountResponse GetAccount(String address) {
        return JSON.parseObject(LiteWallet.INSTANCE.GetAccount(address), GetAccountResponse.class);
    }

    //使用名称查询账户
    public GetAccountResponse GetAccountByName(String name) {
        return JSON.parseObject(LiteWallet.INSTANCE.GetAccountByName(name), GetAccountResponse.class);
    }

    //删除账户
    public DeleteAccountResponse DeleteAccount(String address, String password) {
        return JSON.parseObject(LiteWallet.INSTANCE.DeleteAccount(address, password), DeleteAccountResponse.class);
    }


    //导出账户
    public ExportAccountResponse ExportAccount(String address, String password) {
        return JSON.parseObject(LiteWallet.INSTANCE.ExportAccount(address, password), ExportAccountResponse.class);
    }


    //使用助记词导入账户
    public ImportAccountResponse ImportMnemonic(String mnemonic, String password) {
        return JSON.parseObject(LiteWallet.INSTANCE.ImportMnemonic(mnemonic, password), ImportAccountResponse.class);
    }

    //使用私钥导入账户
    public ImportAccountResponse ImportPrivateKey(String hexPrivateKey, String password) {
        return JSON.parseObject(LiteWallet.INSTANCE.ImportPrivateKey(hexPrivateKey, password), ImportAccountResponse.class);
    }

    //列出所有账户列表
    public ListAccountsResponse ListAllAccounts() {
        return JSON.parseObject(LiteWallet.INSTANCE.ListAllAccounts(), ListAccountsResponse.class);
    }

    //对数据串进行签名
    public SignDataResponse Sign(String address, String password, String signStr) {
        return JSON.parseObject(LiteWallet.INSTANCE.Sign(address, password, signStr), SignDataResponse.class);
    }

    //对base64编码的字符串进行签名
    public SignDataResponse SignBase64(String address, String password, String base64Str) {
        return JSON.parseObject(LiteWallet.INSTANCE.SignBase64(address, password, base64Str), SignDataResponse.class);
    }

    public String getWalletName() {
        return walletName;
    }

    public void setWalletName(String walletName) {
        this.walletName = walletName;
    }

    public String getWalletStoragePath() {
        return walletStoragePath;
    }

    public void setWalletStoragePath(String walletStoragePath) {
        this.walletStoragePath = walletStoragePath;
    }
}


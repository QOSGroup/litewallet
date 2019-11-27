package com.qos.litewallet.demo;

import com.sun.jna.Library;
import com.sun.jna.Native;

/**
 * @author wangzhiyong
 * @date 19-11-27下午2:06
 */

public interface LiteWallet extends Library {


    //Linux下将会查找: CLASSPATH下 liblitewallet.so文件
    //WIN下 TODO
    LiteWallet INSTANCE = (LiteWallet) Native.loadLibrary("litewallet", LiteWallet.class);

    //初始化钱包
    void InitWallet(String name, String storagePath);

    //生成助记词
    String ProduceMnemonic();

    //创建账户
    String CreateAccount(String password);

    //使用指定名称创建账户
    String CreateAccountWithName(String name, String password);

    //使用指定名称和助记词创建账户
    String CreateAccountWithMnemonic(String name, String password, String mnemonic);

    //查询账户
    String GetAccount(String address);

    //使用名称查询账户
    String GetAccountByName(String name);

    //删除账户
    String DeleteAccount(String address, String password);

    //导出账户
    String ExportAccount(String address, String password);

    //使用助记词导入账户
    String ImportMnemonic(String mnemonic, String password);

    //使用私钥导入账户
    String ImportPrivateKey(String hexPrivateKey, String password);

    //列出所有账户列表
    String ListAllAccounts();

    //对数据串进行签名
    String Sign(String address, String password, String signStr);

    //对base64编码的字符串进行签名
    String SignBase64(String address, String password, String base64Str);

}


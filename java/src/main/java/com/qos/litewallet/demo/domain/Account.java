package com.qos.litewallet.demo.domain;

import lombok.Data;

import java.io.Serializable;

/**
 * @author wangzhiyong
 * @date 19-11-27下午3:38
 */

@Data
public class Account implements Serializable {

    /*
     * 账户索引, 从1开始递增
     * */
    private int id;

    /*
     * 账户名称, 名称不能重复, 创建账户未指定名称时, 名称默认为: Account${id}
     * */
    private String name;

    /*
     * QOS账户地址, 地址不可重复
     * */
    private String address;

    /*
     * QOS公钥地址
     * */
    private String publicKey;


    /*
     * 加密之后的私钥串
     * */
    private String privateKeyEnc;

}

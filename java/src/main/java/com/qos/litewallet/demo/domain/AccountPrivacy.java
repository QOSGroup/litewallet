package com.qos.litewallet.demo.domain;

import lombok.Data;

/**
 * @author wangzhiyong
 * @date 19-11-27下午3:42
 */

@Data
public class AccountPrivacy extends Account {

    /*
    * 私钥, 以16进制格式展示
    * */
    private String privateKey;
}

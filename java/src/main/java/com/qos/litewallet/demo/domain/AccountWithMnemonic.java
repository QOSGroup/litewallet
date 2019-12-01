package com.qos.litewallet.demo.domain;

import lombok.Data;

/**
 * @author wangzhiyong
 * @date 19-11-27下午3:38
 */

@Data
public class AccountWithMnemonic extends Account {

    /*
    * 助记词, 12位
    * */
    private String mnemonic;

}

package com.qos.litewallet.demo.response;

import lombok.Data;

/**
 * @author wangzhiyong
 * @date 19-11-27下午3:36
 */

@Data
public class ProduceMnemonicResponse extends BaseResponse {

    /*
    *  助记词
    * */
    private String data;

}

package com.qos.litewallet.demo.response;

import lombok.Data;

/**
 * @author wangzhiyong
 * @date 19-11-27下午4:22
 */

@Data
public class SignDataResponse extends BaseResponse {

    /*
     * 签名之后的字符串(base64编码)
     * */
    private String data;
}

package com.qos.litewallet.demo.response;

import lombok.Data;

import java.io.Serializable;

/**
 * @author wangzhiyong
 * @date 19-11-27下午3:21
 */

@Data
public class BaseResponse implements Serializable {

    //操作结果: 0 成功 1 失败
    protected int code;

    //操作失败时返回的说明
    protected String message;

}

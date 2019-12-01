package com.qos.litewallet.demo.response;

import com.qos.litewallet.demo.domain.Account;
import lombok.Data;

/**
 * @author wangzhiyong
 * @date 19-11-27下午3:54
 */

@Data
public class GetAccountResponse extends BaseResponse {

    private Account data;

}

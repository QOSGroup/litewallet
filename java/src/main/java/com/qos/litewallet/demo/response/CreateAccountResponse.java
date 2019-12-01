package com.qos.litewallet.demo.response;

import com.qos.litewallet.demo.domain.AccountWithMnemonic;
import lombok.Data;

/**
 * @author wangzhiyong
 * @date 19-11-27下午3:44
 */

@Data
public class CreateAccountResponse extends BaseResponse {

    private AccountWithMnemonic data;

}

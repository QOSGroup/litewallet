package com.qos.litewallet.demo.response;

import com.qos.litewallet.demo.domain.Account;
import lombok.Data;

/**
 * @author wangzhiyong
 * @date 19-11-27下午4:10
 */

@Data
public class ImportAccountResponse extends BaseResponse {

    private Account data;

}

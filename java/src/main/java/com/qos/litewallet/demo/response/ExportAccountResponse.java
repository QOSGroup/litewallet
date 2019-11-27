package com.qos.litewallet.demo.response;

import com.qos.litewallet.demo.domain.AccountPrivacy;
import lombok.Data;

/**
 * @author wangzhiyong
 * @date 19-11-27下午3:59
 */

@Data
public class ExportAccountResponse extends BaseResponse {
    private AccountPrivacy data;
}

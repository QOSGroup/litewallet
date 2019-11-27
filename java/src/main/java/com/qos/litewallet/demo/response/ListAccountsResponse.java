package com.qos.litewallet.demo.response;

import com.qos.litewallet.demo.domain.Account;
import lombok.Data;

import java.util.List;

/**
 * @author wangzhiyong
 * @date 19-11-27下午4:07
 */
@Data
public class ListAccountsResponse extends BaseResponse {

    private List<Account> data;

}

package org.exfly.loginweb.file;

import lombok.extern.slf4j.Slf4j;

/**
 * @author zhanghf_mios
 * @description 注销
 * @date 2018年6月26日
 */
@Slf4j
public class LogoutRest {

    /**
     * @author ExFly
     */
    public String execute() {
        try {
            if (log.isDebugEnabled()) {
//				String username = SessionUtil.getAttribute(CommonConstLogin.SESSTION_LOGIN_MARK);
//				log.debug(username+" logout");
            }
//			SessionUtil.clean();
        } catch (Throwable e) {
        }
        return "index";
    }
}

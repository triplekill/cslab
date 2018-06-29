package org.exfly.loginweb.file;

import lombok.extern.slf4j.Slf4j;

import java.io.IOException;
import java.util.Map;
import java.util.Properties;

/**
 * @author zhanghf_mios
 * @description 登陆
 * @date 2018年6月26日
 */
@Slf4j
public class LoginRest {

    private static boolean valid(String username, String password) {
        boolean isvalided = false;

        if (username == null || username.isEmpty() || password == null || password.isEmpty()) {
            return false;
        }

        Properties prop = new Properties();
        try {
            prop.load(LoginRest.class.getClassLoader().getResourceAsStream("config4login.properties"));
            String validPwd = prop.getProperty("login." + username + ".password");
            log.debug(username + " " + validPwd + " " + password);
            isvalided = false;

            if (password.equals(validPwd)) {
                isvalided = true;
            }
//            log.debug(isvalided);
        } catch (IOException e) {
            e.printStackTrace();
            isvalided = false;
        }

        return isvalided;
    }

    public static void main(String[] args) {

        System.out.println(valid("admin", "admin"));
        System.out.println(valid("admin", "admi"));
        System.out.println(valid("admi", "admin"));
        System.out.println(valid("", "admin"));
        System.out.println(valid("admin", ""));
        System.out.println(valid("admin", null));
        System.out.println(valid(null, "admin"));
    }

    /**
     * @description 登陆
     * @author ExFly
     * { "username":"admin",
     * "password":"admin",
     * "randCode":"code"
     * }
     */
    public String execute(Map paramMap) {
//		log.debug(paramMap);
//		log.debug("==========login begin==========");
        //TODO: 参数校验
        String username = ((String) paramMap.get(CommonConstLogin.REQ_JSON_USERNAME)).trim();
        String password = ((String) paramMap.get(CommonConstLogin.REQ_JSON_PASSWORD)).trim();

        //TODO: 验证是否登陆成功，并进行一定的处理
        if (valid(username, password)) {
            //SessionUtil.setAttribute(CommonConstLogin.SESSTION_LOGIN_MARK, username);
        } else {
            //	组织访问
        }
        return "login";
    }
}

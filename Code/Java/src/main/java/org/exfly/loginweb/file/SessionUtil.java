//package org.exfly.loginweb.file;
//
//import javax.servlet.http.HttpSession;
//
//import lombok.extern.log4j.Log4j;
//
//@Log4j
//public class SessionUtil {
//
//	/**
//	 * 校验图形验证码
//	 */
//	public static int checkAuthCode(String code, HttpSession session) {
//		int retCode = 1;
//		// HttpSession session = EchnContextUtils.getEchnContext().getRequest().getSession();
//		if (session != null) {
//			String sessionRand = (String) session.getAttribute("rand");
//			if (sessionRand == null) {
//				retCode = 2;
//			} else {
//				if (sessionRand.equalsIgnoreCase(code)) {
//					retCode = 0;
//				} else {
//					retCode = 1;
//				}
//				session.removeAttribute("rand");
//			}
//		} else {
//			retCode = 2;
//		}
//		return retCode;
//	}
//
//	/**
//	 * 切换渠道
//	 */
//	public static void tabChannel(ChannelInfo channelInfo) {
//		HttpSession session = EchnContextUtils.getEchnContext().getRequest().getSession();
//		if (channelInfo == null) {
//			session.removeAttribute(LocalConstant.CHANNEL_SESSION);
//		} else {
//			session.setAttribute(LocalConstant.CHANNEL_SESSION, channelInfo);
//		}
//	}
//
//	/**
//	 * 获取当前渠道
//	 */
//	public static ChannelInfo getChannel() {
//		HttpSession session = EchnContextUtils.getEchnContext().getRequest().getSession();
//		ChannelInfo channelInfo = null;
//		if (session != null) {
//			channelInfo = (ChannelInfo) session.getAttribute(LocalConstant.CHANNEL_SESSION);
//		}
//		return channelInfo;
//	}
//
//	/**
//	 * @author ExFly
//	 * @return
//	 *
//	 * 获得当前登陆用户名，暗示已经登陆了
//	 */
//	public static String getAttribute(String attr) {
//		HttpSession session = EchnContextUtils.getEchnContext().getRequest().getSession();
//		String username = "";
//		if (session != null) {
//			username = (String) session.getAttribute(attr);
//		}
//		return username;
//	}
//
//	/**
//	 * @author ExFly
//	 * @return
//	 * 设置用户名，暗示正在登陆，此值用来标识已经登陆
//	 */
//	public static String setAttribute(String key, String val) {
//		HttpSession session = EchnContextUtils.getEchnContext().getRequest().getSession();
//		if (session != null) {
//			session.setAttribute(key, val);
//		}
//		return val;
//	}
//
//	public static void  clean() {
//		HttpSession session = EchnContextUtils.getEchnContext().getRequest().getSession();
//		session.invalidate();
//	}
//}

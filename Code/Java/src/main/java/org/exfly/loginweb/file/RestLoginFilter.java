//package org.exfly.loginweb.file;
//
//import java.io.IOException;
//import java.util.HashSet;
//import java.util.Iterator;
//import java.util.Properties;
//import java.util.Set;
//
//
//import lombok.extern.slf4j.Slf4j;
//
//@Slf4j
//public class RestLoginFilter implements Filter {
//
//	private static Set<String> whitelist = new HashSet<>();
//
//	public void destroy() {}
//
//	public void doFilter(ServletRequest request, ServletResponse response,FilterChain filterChain) throws IOException, ServletException {
//		String loginMark = SessionUtil.getAttribute(CommonConstLogin.SESSTION_LOGIN_MARK);
//		HttpServletRequest servletRequest = (HttpServletRequest) request;
//		String path = servletRequest.getRequestURI();
//		HttpServletRequest req = (HttpServletRequest) request;
//		String pathInfo = req.getPathInfo() == null ? "" : req.getPathInfo();
//		String url = req.getServletPath() + pathInfo;
//		try{
//
//			if (RestLoginFilter.contains(url)) {
//				log.info("【REST权限管控】白名单，放行。url=" + url);
//				filterChain.doFilter(request, response);
//				return;
//			}
////			if (path.indexOf("/login")>-1 || path.indexOf("/checkimage.shtml")>-1){
////				filterChain.doFilter(request, response);
////				return;
////			}
//
//			if (loginMark == null || loginMark.isEmpty()) {
//				log.debug("没有登陆Filter");
//				errorResponse(response,null, "未登录");
//
//			}else {
//				filterChain.doFilter(request, response);
//				log.debug("登陆Filter");
//			}
//		}catch (Throwable e){
//			if (!response.isCommitted()) {
//				log.error("【RESTlogin】 " + url + "执行异常（不允许放行），原因：" + e.getMessage(), e);
////				filterChain.doFilter(request, response);
//			}
//		}
//
//	}
//
//	private void errorResponse(ServletResponse response, String url, String msg) throws Exception{
//		response.setContentType("application/json;charset=utf-8");
//		JSONObject json = new JSONObject();
//		json.put("retCode", "-1");
//		json.put("retMsg", msg);
//		json.put("data", "");
//		ServletOutputStream out = response.getOutputStream();
//		out.write(json.toString().getBytes("UTF-8"));
//		out.flush();
//		out.close();
//		response.getOutputStream().close();
//	}
//
//	private static boolean contains(String url) {
//
//		Iterator<String> it = whitelist.iterator();
//		while(it.hasNext()) {
//            String val = it.next();
//            if (val.indexOf(url)>-1){
//        		return true;
//        	}
//        }
//		return false;
////		Properties prop = new Properties();
////        try {
////            prop.load(LoginRest.class.getClassLoader().getResourceAsStream("config4login.properties"));
////            String val;
////            boolean result=false;
////            for(int i=0; (val = prop.getProperty("config4loginwhiteurllist."+i))!=null;i++){
////            	System.out.println(url);
////            	if (val.indexOf(url)>-1){
////            		return true;
////            	}
////            }
////
////        } catch(IOException e) {
////            e.printStackTrace();
////        }
////        return false;
//	}
//
//	public void init(FilterConfig arg0) throws ServletException {
//		Properties prop = new Properties();
//        try {
//            prop.load(LoginRest.class.getClassLoader().getResourceAsStream("config4login.properties"));
//            String val;
//            for(int i=0; (val = prop.getProperty("config4loginwhiteurllist."+i))!=null;i++){
//            	whitelist.add(val);
//            }
//        } catch(IOException e) {
//            e.printStackTrace();
//        }
//	}
//}

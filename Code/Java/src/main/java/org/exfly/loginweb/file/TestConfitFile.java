package org.exfly.loginweb.file;//package com.sitech.echd.gray.rest.login;
//
//import java.io.IOException;
//import java.util.Properties;
//
//public class TestConfitFile {
//	
//	public static void main(String[] args) {
//		reads();
//	}
//	
//	public static void reads() {
//		Properties prop = new Properties();  
//        try { 
//            prop.load(LoginRest.class.getClassLoader().getResourceAsStream("config4login.properties"));
//            String val;
//            boolean result=false;
//            for(int i=0; (val = prop.getProperty("config4loginwhiteurllist."+i))!=null;i++){
//            	String url = prop.getProperty("config4loginwhiteurllist."+i);
//            	System.out.println(url);
//            	if (url.indexOf("/login")>-1){
//            		result=true;
//            	}
//            	System.out.println(result);
//            	result = false;
//            }
//        } catch(IOException e) {
//            e.printStackTrace();
//        }
//	}
//}

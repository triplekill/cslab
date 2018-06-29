package org.exfly.io;

import java.io.IOException;
import java.net.URI;
import java.net.URISyntaxException;
import java.nio.charset.Charset;
import java.nio.file.Files;
import java.nio.file.Paths;
import java.util.List;


public class FileProcess {

    /**
     * @param filename
     * @return 文件中所有字符
     */
    private static String readAllStringInFile(String filename) {
        String result = "";
        try {
            StringBuilder sb = new StringBuilder();
            Charset charset = Charset.forName("UTF-8");
            URI uri = ClassLoader.getSystemResource(filename).toURI();
            String mainPath = Paths.get(uri).toString();
//			log.debug(mainPath);
            java.nio.file.Path textFiles = Paths.get(mainPath);
            List<String> linesRead = null;
            try {
                linesRead = Files.readAllLines(textFiles, charset);
            } catch (IOException e) {
                e.printStackTrace();
            }
            if (linesRead != null) {
                for (String line : linesRead) {
                    sb.append(line);
                }
            }
            result = sb.toString();
        } catch (URISyntaxException e) {
            e.printStackTrace();
            result = "";
        }
        return result;

    }

    public static void main(String[] args) {
//		String p = "D:/Project/shixi/gray-login/gray-rest/src/main/resources/lua/gray_config.json";
//		String p = RestGetAllConfigJson.class.getClassLoader().getResource("/lua/gray_config.json").getPath();
//		String p = Thread.currentThread().getContextClassLoader().getResource("lua/gray_config.json").getFile();
//		String p = RestGetAllConfigJson.class.getResource("/").getFile()+"lua/gray_config.json";
        System.out.println(readAllStringInFile("lua/gray_config.json"));
    }
}


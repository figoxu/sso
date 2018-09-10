package me.figoxu.sso.web;


import org.apache.log4j.Logger;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestMethod;
import org.springframework.web.bind.annotation.RestController;
import org.springframework.web.context.request.RequestContextHolder;
import org.springframework.web.context.request.ServletRequestAttributes;

import javax.servlet.http.HttpServletResponse;
import java.io.PrintWriter;

@RestController
@RequestMapping(value = "/main")
public class MainController {
    Logger logger = Logger.getLogger(MainController.class);


    //http://jsp.localhost/main/welcome
    @RequestMapping(value = "/welcome", method = RequestMethod.GET)
    public ResponseEntity main() throws Exception {
        HttpServletResponse response = ((ServletRequestAttributes)RequestContextHolder.getRequestAttributes()).getResponse();
        logger.info("成功进入页面");
        PrintWriter writer = response.getWriter();
        writer.write("<h1>Welcome To Java Page</h1>");
        writer.close();
        return new ResponseEntity(HttpStatus.OK);
    }
}

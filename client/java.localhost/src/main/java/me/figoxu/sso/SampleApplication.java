package me.figoxu.sso;

import me.figoxu.sso.filter.SsoFilter;
import me.figoxu.sso.util.SpringContextUtil;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.boot.web.servlet.FilterRegistrationBean;
import org.springframework.boot.web.servlet.ServletComponentScan;
import org.springframework.context.ConfigurableApplicationContext;
import org.springframework.context.annotation.Bean;

@SpringBootApplication
@ServletComponentScan
public class SampleApplication {
    public static void main(String[] args) {
        ConfigurableApplicationContext context = SpringApplication.run(SampleApplication.class, args);
        SpringContextUtil.setApplicationContext(context);
    }


    @Bean
    public FilterRegistrationBean ssoFilterRegistration() {
        FilterRegistrationBean registration = new FilterRegistrationBean();
        registration.setFilter(new SsoFilter());
        registration.addUrlPatterns("/*");
        registration.addInitParameter("paramName", "paramValue");
        registration.setName("ssoFilter");
        registration.setOrder(1);
        return registration;
    }
}

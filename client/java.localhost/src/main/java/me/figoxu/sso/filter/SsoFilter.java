package me.figoxu.sso.filter;

import io.grpc.ManagedChannel;
import io.grpc.ManagedChannelBuilder;
import me.figoxu.sso.cache.SsoCache;
import me.figoxu.sso.grpc.Api;
import me.figoxu.sso.grpc.SsoServiceGrpc;
import me.figoxu.sso.util.SpringContextUtil;
import me.figoxu.sso.web.MainController;
import org.apache.log4j.Logger;

import javax.servlet.*;
import javax.servlet.http.Cookie;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;
import java.io.IOException;
import java.net.URLEncoder;

/**
 * Created by xujianhui on 2018/9/9.
 */
public class SsoFilter implements Filter {
    Logger logger = Logger.getLogger(SsoFilter.class);

    public static final String redirect_url = "http://sso.localhost/sso/login/redirect";
    public static final String SSO_BASIC_PURE_TOKEN = "basic_raw_token";
    public static final String SSO_TOKEN_COOKIE = "sso";
    public static final String SSO_TOKEN_PARAM = "basic_pure_token";
    public static final String SSO_FROM_PARAM = "from";
    public static final String MINE_DOMAIN = "java.localhost";
    public static final String GRPC_HOST = "127.0.0.1";
    public static final Integer GRPC_PORT = 8084;


    @Override
    public void init(FilterConfig filterConfig) throws ServletException {

    }

    private void resolveRedirectUrl(HttpServletRequest req, HttpServletResponse rsp) throws IOException {
        String currentURL = req.getRequestURL().toString();
        String queryString = req.getQueryString();
        if (queryString != null && queryString.length() > 0) {
            currentURL = currentURL + "?" + queryString;
        }
        logger.info(currentURL);
        logger.info(URLEncoder.encode(currentURL, "UTF-8"));
        String redirectURL = redirect_url + "?from=" + URLEncoder.encode(currentURL, "UTF-8");
        rsp.sendRedirect(redirectURL);
    }

    private void autoSaveToken(HttpServletRequest req, HttpServletResponse rsp) throws IOException {
        String token = req.getParameter(SSO_TOKEN_PARAM);
        if (token == null || token.length() <= 0) {
            return;
        }
        Cookie cookie = new Cookie(SSO_TOKEN_COOKIE, token);
        cookie.setPath("/");
        cookie.setMaxAge(60 * 60 * 24);
        cookie.setDomain(MINE_DOMAIN);
        rsp.addCookie(cookie);
        req.getSession().setAttribute(SSO_BASIC_PURE_TOKEN, token);
    }

    private String getBasicPureToken(HttpServletRequest req, HttpServletResponse rsp) throws IOException {
        String token = req.getParameter(SSO_TOKEN_PARAM);
        if (token != null && token.length() > 0) {
            System.out.println("GET TOKEN FROM REQUEST");
            return token;
        }
        Cookie[] cookies = req.getCookies();
        System.out.println(cookies);
        if (cookies != null) {
            for (Cookie cookie : cookies) {
                if (cookie.getName().equals(SSO_TOKEN_COOKIE)) {
                    token = cookie.getValue();
                    if (token != null && token.length() > 0) {
                        System.out.println("GET TOKEN FROM COOKIE");
                        return token;
                    }
                }
            }
        }
        token = (String) req.getSession().getAttribute(SSO_BASIC_PURE_TOKEN);
        return token;
    }

    private Boolean isLogin(HttpServletRequest req, HttpServletResponse rsp) throws IOException {
        String token = getBasicPureToken(req, rsp);
        if (token == null || token.length() <= 0) {
            return false;
        }
        SsoCache ssoCache = SpringContextUtil.getBean(SsoCache.class);
        Api.User user = ssoCache.getUser(token);
        if (user == null || user.getId() <= 0) {
            return false;
        }
        return true;
    }

    @Override
    public void doFilter(ServletRequest servletRequest, ServletResponse servletResponse, FilterChain filterChain) throws IOException, ServletException {
        HttpServletRequest req = (HttpServletRequest) servletRequest;
        HttpServletResponse rsp = (HttpServletResponse) servletResponse;
        autoSaveToken(req, rsp);
        if (!isLogin(req, rsp)) {
            resolveRedirectUrl(req, rsp);
        } else {
            filterChain.doFilter(servletRequest, servletResponse);
        }
    }

    @Override
    public void destroy() {

    }


}

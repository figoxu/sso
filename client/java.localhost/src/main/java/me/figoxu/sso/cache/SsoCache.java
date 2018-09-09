package me.figoxu.sso.cache;

import io.grpc.ManagedChannel;
import io.grpc.ManagedChannelBuilder;
import me.figoxu.sso.filter.SsoFilter;
import me.figoxu.sso.grpc.Api;
import me.figoxu.sso.grpc.SsoServiceGrpc;
import org.springframework.stereotype.Component;
import org.springframework.cache.annotation.Cacheable;

@Component
public class SsoCache {
    @Cacheable(value = "sso_cache", key = "'user_token_'+#token")
    public Api.User getUser(String token) {
        System.out.println(">>>>通过GRPC获取User信息");
        ManagedChannel channel = ManagedChannelBuilder.forAddress(SsoFilter.GRPC_HOST, SsoFilter.GRPC_PORT)
                .usePlaintext(true)
                .build();
        SsoServiceGrpc.SsoServiceBlockingStub ssoServiceBlockingStub = SsoServiceGrpc.newBlockingStub(channel);
        Api.LoginInfoReq loginInfoReq = Api.LoginInfoReq.newBuilder().setBasicRawToken("").build();
        Api.LoginInfoRsp resp = ssoServiceBlockingStub.getLoginInfo(loginInfoReq);
        Api.User user = resp.getUser();
        return user;
    }
}

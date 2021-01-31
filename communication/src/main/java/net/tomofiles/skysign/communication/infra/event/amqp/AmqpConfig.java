package net.tomofiles.skysign.communication.infra.event.amqp;

import org.springframework.amqp.core.FanoutExchange;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;

@Configuration
public class AmqpConfig {

    @Value("${skysign.event.communication_id_changed_event}")
    private String EXCHANGE_NAME;

	@Bean
	public FanoutExchange exchange() {
	  	return new FanoutExchange(EXCHANGE_NAME, false, true);
	}

}
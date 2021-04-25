package net.tomofiles.skysign.communication.infra.communication;

import java.util.List;

import org.apache.ibatis.annotations.Insert;
import org.apache.ibatis.annotations.Mapper;
import org.apache.ibatis.annotations.Select;
import org.apache.ibatis.annotations.Update;

@Mapper
public interface CommandMapper {
    @Select("SELECT id, communication_id as communicationId, type, time FROM commands WHERE communication_id = #{communicationId}")
    List<CommandRecord> findByCommunicationId(String communicationId);

    @Insert("INSERT INTO commands (id, communication_id, type, time) VALUES (#{id}, #{communicationId}, #{type}, #{time})")
    void create(CommandRecord command);

    @Update("DELETE FROM commands WHERE id = #{id}")
    void delete(String id);

    @Update("DELETE FROM commands WHERE communication_id = #{communicationId}")
    void deleteByCommunicationId(String communicationId);
}
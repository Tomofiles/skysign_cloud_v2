package net.tomofiles.skysign.communication.infra.communication;

import java.util.List;

import org.apache.ibatis.annotations.Insert;
import org.apache.ibatis.annotations.Mapper;
import org.apache.ibatis.annotations.Select;
import org.apache.ibatis.annotations.Update;

@Mapper
public interface CommandMapper {
    @Select("SELECT id, comm_id as commId, type, time FROM command WHERE comm_id = #{commId} ORDER BY time")
    List<CommandRecord> findByCommId(String commId);

    @Insert("INSERT INTO command (id, comm_id, type, time) VALUES (#{id}, #{commId}, #{type}, #{time})")
    void create(CommandRecord command);

    @Update("DELETE FROM command WHERE id = #{id}")
    void delete(String id);

    @Update("DELETE FROM command WHERE comm_id = #{commId}")
    void deleteByCommId(String commId);
}
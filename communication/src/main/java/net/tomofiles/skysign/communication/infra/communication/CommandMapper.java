package net.tomofiles.skysign.communication.infra.communication;

import java.util.List;

import org.apache.ibatis.annotations.Insert;
import org.apache.ibatis.annotations.Mapper;
import org.apache.ibatis.annotations.Select;
import org.apache.ibatis.annotations.Update;

@Mapper
public interface CommandMapper {
    @Select("SELECT id, comm_id as commId, type FROM command WHERE comm_id = #{commId}")
    List<CommandRecord> findByCommId(String commId);

    @Insert("INSERT INTO command (id, comm_id, type) VALUES (#{id}, #{commId}, #{type})")
    void create(CommandRecord command);

    @Update("DELETE FROM command WHERE id = #{id}")
    void delete(String id);

    @Update("DELETE FROM command WHERE comm_id = #{commId}")
    void deleteByCommId(String commId);
}
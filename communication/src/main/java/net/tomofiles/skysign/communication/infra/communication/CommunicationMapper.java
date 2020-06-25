package net.tomofiles.skysign.communication.infra.communication;

import org.apache.ibatis.annotations.Insert;
import org.apache.ibatis.annotations.Mapper;
import org.apache.ibatis.annotations.Select;
import org.apache.ibatis.annotations.Update;

@Mapper
public interface CommunicationMapper {
    @Select("SELECT id, mission_id as missionId, version FROM communication WHERE id = #{id}")
    CommunicationRecord find(String id);

    @Insert("INSERT INTO communication (id, mission_id, version) VALUES (#{id}, #{missionId}, #{version})")
    void create(CommunicationRecord communication);

    @Update("UPDATE communication SET mission_id = #{missionId}, version = #{version} + 1 WHERE id = #{id} and version = #{version}")
    void update(CommunicationRecord communication);

    @Update("DELETE FROM communication WHERE id = #{id} and version = #{version}")
    void delete(DeleteCondition condition);
}
package net.tomofiles.skysign.communication.infra.communication;

import org.apache.ibatis.annotations.Insert;
import org.apache.ibatis.annotations.Mapper;
import org.apache.ibatis.annotations.Select;
import org.apache.ibatis.annotations.Update;

@Mapper
public interface CommunicationMapper {
    @Select("SELECT id, mission_id as missionId FROM communication WHERE id = #{id} FOR UPDATE")
    CommunicationRecord find(String id);

    @Insert("INSERT INTO communication (id, mission_id) VALUES (#{id}, #{missionId})")
    void create(CommunicationRecord communication);

    @Update("UPDATE communication SET mission_id = #{missionId} WHERE id = #{id}")
    void update(CommunicationRecord communication);

    @Update("DELETE FROM communication WHERE id = #{id}")
    void delete(String id);
}
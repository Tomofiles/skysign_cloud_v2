package net.tomofiles.skysign.communication.infra.communication;

import java.util.List;

import org.apache.ibatis.annotations.Insert;
import org.apache.ibatis.annotations.Mapper;
import org.apache.ibatis.annotations.Select;
import org.apache.ibatis.annotations.Update;

@Mapper
public interface CommunicationMapper {
    @Select("SELECT id, vehicle_id as vehicleId, controlled, mission_id as missionId FROM communication")
    List<CommunicationRecord> findAll();

    @Select("SELECT id, vehicle_id as vehicleId, controlled, mission_id as missionId FROM communication WHERE id = #{id} FOR UPDATE")
    CommunicationRecord find(String id);

    @Insert("INSERT INTO communication (id, vehicle_id, controlled, mission_id) VALUES (#{id}, #{vehicleId}, #{controlled}, #{missionId})")
    void create(CommunicationRecord communication);

    @Update("UPDATE communication SET vehicle_id = #{vehicleId}, controlled = #{controlled}, mission_id = #{missionId} WHERE id = #{id}")
    void update(CommunicationRecord communication);

    @Update("DELETE FROM communication WHERE id = #{id}")
    void delete(String id);
}
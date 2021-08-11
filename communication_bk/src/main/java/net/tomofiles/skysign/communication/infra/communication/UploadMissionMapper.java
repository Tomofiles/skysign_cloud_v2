package net.tomofiles.skysign.communication.infra.communication;

import java.util.List;

import org.apache.ibatis.annotations.Insert;
import org.apache.ibatis.annotations.Mapper;
import org.apache.ibatis.annotations.Select;
import org.apache.ibatis.annotations.Update;

@Mapper
public interface UploadMissionMapper {
    @Select("SELECT id, communication_id as communicationId, mission_id as missionId FROM upload_missions WHERE communication_id = #{communicationId}")
    List<UploadMissionRecord> findByCommunicationId(String communicationId);

    @Insert("INSERT INTO upload_missions (id, communication_id, mission_id) VALUES (#{id}, #{communicationId}, #{missionId})")
    void create(UploadMissionRecord command);

    @Update("DELETE FROM upload_missions WHERE id = #{id}")
    void delete(String id);

    @Update("DELETE FROM upload_missions WHERE communication_id = #{communicationId}")
    void deleteByCommunicationId(String communicationId);
}
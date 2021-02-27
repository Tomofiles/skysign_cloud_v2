package net.tomofiles.skysign.communication.infra.communication;

import java.util.List;

import org.apache.ibatis.annotations.Insert;
import org.apache.ibatis.annotations.Mapper;
import org.apache.ibatis.annotations.Select;
import org.apache.ibatis.annotations.Update;

@Mapper
public interface UploadMissionMapper {
    @Select("SELECT id, comm_id as commId, mission_id as missionId FROM upload_mission WHERE comm_id = #{commId}")
    List<UploadMissionRecord> findByCommId(String commId);

    @Insert("INSERT INTO upload_mission (id, comm_id, mission_id) VALUES (#{id}, #{commId}, #{missionId})")
    void create(UploadMissionRecord command);

    @Update("DELETE FROM upload_mission WHERE id = #{id}")
    void delete(String id);

    @Update("DELETE FROM upload_mission WHERE comm_id = #{commId}")
    void deleteByCommId(String commId);
}
package net.tomofiles.skysign.mission.infra.mission;

import java.util.List;

import org.apache.ibatis.annotations.Insert;
import org.apache.ibatis.annotations.Mapper;
import org.apache.ibatis.annotations.Select;
import org.apache.ibatis.annotations.Update;

@Mapper
public interface WaypointMapper {
    @Select("SELECT "
            + "mission_id as missionId, "
            + "order, "
            + "latitude, "
            + "longitude, "
            + "height_wgs84_ellipsoid_m as heightWGS84EllipsoidM, "
            + "speed_m_s as speedMS "
            + "FROM waypoint WHERE mission_id = #{id}")
    List<WaypointRecord> find(String id);

    @Insert("INSERT INTO waypoint "
            + "(mission_id, order, latitude, longitude, height_wgs84_ellipsoid_m, speed_m_s) "
            + "VALUES (#{missionId}, #{order}, #{latitude}, #{longitude}, #{heightWGS84EllipsoidM}, #{speedMS})")
    void create(WaypointRecord waypoint);

    @Update("DELETE FROM waypoint WHERE mission_id = #{id}")
    void delete(String id);
}
package net.tomofiles.skysign.communication.infra.mission;

import java.util.List;

import org.apache.ibatis.annotations.Insert;
import org.apache.ibatis.annotations.Mapper;
import org.apache.ibatis.annotations.Select;
import org.apache.ibatis.annotations.Update;

@Mapper
public interface WaypointMapper {
    @Select("SELECT "
            + "mission_id as missionId, "
            + "point_order as order, "
            + "latitude_degree as latitudeDegree, "
            + "longitude_degree as longitudeDegree, "
            + "relative_height_m as relativeHeightM, "
            + "speed_ms as speedMS "
            + "FROM waypoints WHERE mission_id = #{id}")
    List<WaypointRecord> find(String id);

    @Insert("INSERT INTO waypoints "
            + "(mission_id, point_order, latitude_degree, longitude_degree, relative_height_m, speed_ms) "
            + "VALUES (#{missionId}, #{order}, #{latitudeDegree}, #{longitudeDegree}, #{relativeHeightM}, #{speedMS})")
    void create(WaypointRecord waypoint);

    @Update("DELETE FROM waypoints WHERE mission_id = #{id}")
    void delete(String id);
}
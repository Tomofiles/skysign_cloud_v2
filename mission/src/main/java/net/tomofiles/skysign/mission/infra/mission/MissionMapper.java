package net.tomofiles.skysign.mission.infra.mission;

import java.util.List;

import org.apache.ibatis.annotations.Insert;
import org.apache.ibatis.annotations.Mapper;
import org.apache.ibatis.annotations.Select;
import org.apache.ibatis.annotations.Update;

import net.tomofiles.skysign.mission.infra.common.DeleteCondition;

@Mapper
public interface MissionMapper {
    @Select("SELECT "
            + "id, "
            + "name, "
            + "takeoff_point_ground_height_wgs84_ellipsoid_m as takeoffPointGroundHeightWGS84EllipsoidM, "
            + "is_carbon_copy as isCarbonCopy, "
            + "version "
            + "FROM mission WHERE id = #{id}")
    MissionRecord find(String id);

    @Select("SELECT "
            + "id, "
            + "name, "
            + "takeoff_point_ground_height_wgs84_ellipsoid_m as takeoffPointGroundHeightWGS84EllipsoidM, "
            + "is_carbon_copy as isCarbonCopy, "
            + "version "
            + "FROM mission")
    List<MissionRecord> findAll();

    @Select("SELECT "
            + "id, "
            + "name, "
            + "takeoff_point_ground_height_wgs84_ellipsoid_m as takeoffPointGroundHeightWGS84EllipsoidM, "
            + "is_carbon_copy as isCarbonCopy, "
            + "version "
            + "FROM mission "
            + "WHERE is_carbon_copy = false")
    List<MissionRecord> findAllOriginal();

    @Insert("INSERT INTO mission "
            + "(id, name, takeoff_point_ground_height_wgs84_ellipsoid_m, is_carbon_copy, version) "
            + "VALUES (#{id}, #{name}, #{takeoffPointGroundHeightWGS84EllipsoidM}, #{isCarbonCopy}, #{version})")
    void create(MissionRecord mission);

    @Update("UPDATE mission SET "
            + "name = #{name}, "
            + "takeoff_point_ground_height_wgs84_ellipsoid_m = #{takeoffPointGroundHeightWGS84EllipsoidM}, "
            + "is_carbon_copy = #{isCarbonCopy}, "
            + "version = #{newVersion} "
            + "WHERE id = #{id} AND version = #{version}")
    void update(MissionRecord mission);

    @Update("DELETE FROM mission WHERE id = #{id} AND version = #{version}")
    void delete(DeleteCondition condition);
}
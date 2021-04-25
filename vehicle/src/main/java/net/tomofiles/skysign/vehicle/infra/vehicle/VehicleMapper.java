package net.tomofiles.skysign.vehicle.infra.vehicle;

import java.util.List;

import org.apache.ibatis.annotations.Insert;
import org.apache.ibatis.annotations.Mapper;
import org.apache.ibatis.annotations.Select;
import org.apache.ibatis.annotations.Update;

import net.tomofiles.skysign.vehicle.infra.common.DeleteCondition;

@Mapper
public interface VehicleMapper {
    @Select("SELECT id, name, communication_id as communicationId, is_carbon_copy as isCarbonCopy, version FROM vehicles WHERE id = #{id}")
    VehicleRecord find(String id);

    @Select("SELECT id, name, communication_id as communicationId, is_carbon_copy as isCarbonCopy, version FROM vehicles")
    List<VehicleRecord> findAll();

    @Select("SELECT id, name, communication_id as communicationId, is_carbon_copy as isCarbonCopy, version FROM vehicles WHERE is_carbon_copy = false")
    List<VehicleRecord> findAllOriginal();

    @Insert("INSERT INTO vehicles (id, name, communication_id, is_carbon_copy, version) VALUES (#{id}, #{name}, #{communicationId}, #{isCarbonCopy}, #{version})")
    void create(VehicleRecord vehicle);

    @Update("UPDATE vehicles SET name = #{name}, communication_id = #{communicationId}, is_carbon_copy = #{isCarbonCopy}, version = #{newVersion} WHERE id = #{id} AND version = #{version}")
    void update(VehicleRecord vehicle);

    @Update("DELETE FROM vehicles WHERE id = #{id} AND version = #{version}")
    void delete(DeleteCondition condition);
}
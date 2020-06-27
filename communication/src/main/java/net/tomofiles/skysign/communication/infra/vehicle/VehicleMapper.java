package net.tomofiles.skysign.communication.infra.vehicle;

import java.util.List;

import org.apache.ibatis.annotations.Insert;
import org.apache.ibatis.annotations.Mapper;
import org.apache.ibatis.annotations.Select;
import org.apache.ibatis.annotations.Update;

import net.tomofiles.skysign.communication.infra.common.DeleteCondition;

@Mapper
public interface VehicleMapper {
    @Select("SELECT id, name, comm_id as commId, version FROM vehicle WHERE id = #{id}")
    VehicleRecord find(String id);

    @Select("SELECT id, name, comm_id as commId, version FROM vehicle")
    List<VehicleRecord> findAll();

    @Insert("INSERT INTO vehicle (id, name, comm_id, version) VALUES (#{id}, #{name}, #{commId}, #{version})")
    void create(VehicleRecord vehicle);

    @Update("UPDATE vehicle SET name = #{name}, comm_id = #{commId}, version = #{version} + 1 WHERE id = #{id} AND version = #{version}")
    void update(VehicleRecord vehicle);

    @Update("DELETE FROM vehicle WHERE id = #{id} AND version = #{version}")
    void delete(DeleteCondition condition);
}
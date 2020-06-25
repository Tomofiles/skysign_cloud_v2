package net.tomofiles.skysign.communication.infra;

import java.util.List;

import org.apache.ibatis.annotations.Insert;
import org.apache.ibatis.annotations.Mapper;
import org.apache.ibatis.annotations.Select;

@Mapper
public interface VehicleRepository {
    @Select("SELECT id, name FROM vehicle")
    List<Vehicle> findAll();

    @Insert("INSERT INTO vehicle (id, name) VALUES (#{id}, #{name})")
    void create(Vehicle vehicle);
}
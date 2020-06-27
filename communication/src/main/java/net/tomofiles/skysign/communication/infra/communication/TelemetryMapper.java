package net.tomofiles.skysign.communication.infra.communication;

import org.apache.ibatis.annotations.Insert;
import org.apache.ibatis.annotations.Mapper;
import org.apache.ibatis.annotations.Select;
import org.apache.ibatis.annotations.Update;

@Mapper
public interface TelemetryMapper {
    @Select("SELECT"
        + " comm_id as commId,"
        + " latitude,"
        + " longitude,"
        + " altitude,"
        + " speed,"
        + " armed,"
        + " flight_mode as flightMode,"
        + " orientation_x as oriX,"
        + " orientation_y as oriY,"
        + " orientation_z as oriZ,"
        + " orientation_w as oriW"
        + " FROM telemetry" 
        + " WHERE comm_id = #{commId}")
    TelemetryRecord find(String id);

    @Insert("INSERT INTO telemetry ("
        + " comm_id,"
        + " latitude,"
        + " longitude,"
        + " altitude,"
        + " speed,"
        + " armed,"
        + " flight_mode,"
        + " orientation_x,"
        + " orientation_y,"
        + " orientation_z,"
        + " orientation_w"
        + ") VALUES ("
        + " #{commId},"
        + " #{latitude},"
        + " #{longitude},"
        + " #{altitude},"
        + " #{speed},"
        + " #{armed},"
        + " #{flightMode},"
        + " #{oriX},"
        + " #{oriY},"
        + " #{oriZ},"
        + " #{oriW}"
        + ")")
    void create(TelemetryRecord communication);

    @Update("UPDATE telemetry SET"
        + " latitude = #{latitude},"
        + " longitude = #{longitude},"
        + " altitude = #{altitude},"
        + " speed = #{speed},"
        + " armed = #{armed},"
        + " flight_mode = #{flightMode},"
        + " orientation_x = #{oriX},"
        + " orientation_y = #{oriY},"
        + " orientation_z = #{oriZ},"
        + " orientation_w = #{oriW}"
        + " WHERE"
        + " comm_id = #{commId}")
    void update(TelemetryRecord communication);

    @Update("DELETE FROM telemetry WHERE id = #{id}")
    void delete(String id);
}


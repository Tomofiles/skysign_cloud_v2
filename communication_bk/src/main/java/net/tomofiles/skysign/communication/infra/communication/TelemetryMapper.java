package net.tomofiles.skysign.communication.infra.communication;

import org.apache.ibatis.annotations.Insert;
import org.apache.ibatis.annotations.Mapper;
import org.apache.ibatis.annotations.Select;
import org.apache.ibatis.annotations.Update;

@Mapper
public interface TelemetryMapper {
    @Select("SELECT"
        + " communication_id as communicationId,"
        + " latitude,"
        + " longitude,"
        + " altitude,"
        + " relative_altitude as relativeAltitude,"
        + " speed,"
        + " armed,"
        + " flight_mode as flightMode,"
        + " orientation_x as oriX,"
        + " orientation_y as oriY,"
        + " orientation_z as oriZ,"
        + " orientation_w as oriW"
        + " FROM telemetries" 
        + " WHERE communication_id = #{communicationId}")
    TelemetryRecord find(String communicationId);

    @Insert("INSERT INTO telemetries ("
        + " communication_id,"
        + " latitude,"
        + " longitude,"
        + " altitude,"
        + " relative_altitude,"
        + " speed,"
        + " armed,"
        + " flight_mode,"
        + " orientation_x,"
        + " orientation_y,"
        + " orientation_z,"
        + " orientation_w"
        + ") VALUES ("
        + " #{communicationId},"
        + " #{latitude},"
        + " #{longitude},"
        + " #{altitude},"
        + " #{relativeAltitude},"
        + " #{speed},"
        + " #{armed},"
        + " #{flightMode},"
        + " #{oriX},"
        + " #{oriY},"
        + " #{oriZ},"
        + " #{oriW}"
        + ")")
    void create(TelemetryRecord communication);

    @Update("UPDATE telemetries SET"
        + " latitude = #{latitude},"
        + " longitude = #{longitude},"
        + " altitude = #{altitude},"
        + " relative_altitude = #{relativeAltitude},"
        + " speed = #{speed},"
        + " armed = #{armed},"
        + " flight_mode = #{flightMode},"
        + " orientation_x = #{oriX},"
        + " orientation_y = #{oriY},"
        + " orientation_z = #{oriZ},"
        + " orientation_w = #{oriW}"
        + " WHERE"
        + " communication_id = #{communicationId}")
    void update(TelemetryRecord communication);

    @Update("DELETE FROM telemetries WHERE communication_id = #{communicationId}")
    void delete(String communicationId);
}


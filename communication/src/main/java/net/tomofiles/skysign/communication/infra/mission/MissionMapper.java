package net.tomofiles.skysign.communication.infra.mission;

import org.apache.ibatis.annotations.Insert;
import org.apache.ibatis.annotations.Mapper;
import org.apache.ibatis.annotations.Select;
import org.apache.ibatis.annotations.Update;

@Mapper
public interface MissionMapper {
    @Select("SELECT "
            + "id "
            + "FROM missions WHERE id = #{id}")
    MissionRecord find(String id);

    @Insert("INSERT INTO missions "
            + "(id) "
            + "VALUES (#{id})")
    void create(MissionRecord mission);

    @Update("DELETE FROM missions WHERE id = #{id}")
    void delete(String id);
}
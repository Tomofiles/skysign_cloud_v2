package net.tomofiles.skysign.communication.infra.communication;

import java.util.List;

import org.apache.ibatis.annotations.Insert;
import org.apache.ibatis.annotations.Mapper;
import org.apache.ibatis.annotations.Select;
import org.apache.ibatis.annotations.Update;

@Mapper
public interface CommunicationMapper {
    @Select("SELECT id, controlled FROM communication")
    List<CommunicationRecord> findAll();

    @Select("SELECT id, controlled FROM communication WHERE id = #{id} FOR UPDATE")
    CommunicationRecord find(String id);

    @Insert("INSERT INTO communication (id, controlled) VALUES (#{id}, #{controlled})")
    void create(CommunicationRecord communication);

    @Update("UPDATE communication SET controlled = #{controlled} WHERE id = #{id}")
    void update(CommunicationRecord communication);

    @Update("DELETE FROM communication WHERE id = #{id}")
    void delete(String id);
}
package net.tomofiles.skysign.communication.infra.communication;

import java.util.List;

import org.apache.ibatis.annotations.Insert;
import org.apache.ibatis.annotations.Mapper;
import org.apache.ibatis.annotations.Select;
import org.apache.ibatis.annotations.Update;

@Mapper
public interface CommunicationMapper {
    @Select("SELECT id FROM communications")
    List<CommunicationRecord> findAll();

    @Select("SELECT id FROM communications WHERE id = #{id} FOR UPDATE")
    CommunicationRecord find(String id);

    @Insert("INSERT INTO communications (id) VALUES (#{id})")
    void create(CommunicationRecord communication);

    @Update("DELETE FROM communications WHERE id = #{id}")
    void delete(String id);
}
package net.tomofiles.skysign.mission.domain.mission;

import org.junit.jupiter.api.Test;

import static com.google.common.truth.Truth.assertThat;

import static net.tomofiles.skysign.mission.domain.mission.MissionObjectMother.newSingleNavigation;
import static net.tomofiles.skysign.mission.domain.mission.MissionObjectMother.newSeveralNavigation;
import static net.tomofiles.skysign.mission.domain.mission.MissionObjectMother.newSeveralInRondomOrderNavigation;;

public class MissionNavigationTests {
    
    /**
     * Mission Navigationの1件同士の比較結果が同一であること。
     */
    @Test
    public void sameTwoMissionNavigationInSingleTest() {
        assertThat(newSingleNavigation()).isEqualTo(newSingleNavigation());
    }
    
    /**
     * Mission Navigationの複数件かつ同順序同士の比較結果が同一であること。
     */
    @Test
    public void sameTwoMissionNavigationInSeveralInSameOrderTest() {
        assertThat(newSeveralNavigation()).isEqualTo(newSeveralNavigation());
    }
    
    /**
     * Mission Navigationの複数件かつ異なる順序同士の比較結果が同一でないこと。
     */
    @Test
    public void differentTwoMissionNavigationInSeveralInAnotherOrderTest() {
        assertThat(newSeveralNavigation()).isNotEqualTo(newSeveralInRondomOrderNavigation());
    }
}
<template>
    <EditableLabelSelect
            :edit-mode="editMode"
            :value="team.name"
            :options="teams"
            :callback="updateTeamName"
    />
</template>

<script>
    import EditableLabelSelect from "../../common/EditableLabelSelect";
    import {submitChange} from "../../../submit";

    export const TEAMS = [
        "",
        "ZJUNlict",
        "SRC赛队",
        "编译错误",
        "StepForwards",
        "简单难度的电脑",
        "江苏海洋大学1队",
        "黄河1号",
        "CUGB",
        "CUGB2",
        "空指针异常",
        "ZJUTKnights",
        "请求超时",
    ];

    export default {
        name: "TeamName",
        components: {EditableLabelSelect},
        props: {
            teamColor: String,
            editMode: Object,
        },
        computed: {
            team: function () {
                return this.$store.state.matchState.teamState[this.teamColor]
            },
            teams: function () {
                return TEAMS
            }
        },
        methods: {
            updateTeamName: function (v) {
                submitChange({
                    updateTeamState: {
                        forTeam: this.teamColor,
                        teamName: v
                    }
                });
            }
        }
    }
</script>

<style scoped>

</style>

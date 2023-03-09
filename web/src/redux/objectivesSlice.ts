import { UserInfo } from "@/types/shared";
import { createSlice, PayloadAction } from "@reduxjs/toolkit";
import moment from "moment";
import { persistReducer } from 'redux-persist';
import storage from 'redux-persist/lib/storage';
import { Objective } from "@/types/shared";

interface ObjectiveState {
    errors?: string[];
    loading: boolean;
    objectives: Objective[]
}

const initialState: ObjectiveState = {
    loading: false,
    objectives: []
};

const objectiveslice = createSlice({
    name: 'okr',
    initialState,
    reducers: {
        loadObjectivesRequest(state, action) {
            state.objectives = action.payload;
        },
        loadObjectivesSuccess: (state, { payload: { data } }) => {
            state.objectives = data;
        }
    }
});

export const objectiveAction = objectiveslice.actions;
export default objectiveslice.reducer;
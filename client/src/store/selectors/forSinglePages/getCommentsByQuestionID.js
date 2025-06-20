import { createSelector } from 'reselect';

const getComments = state => state.main.comments
const getQuestionID = (_, id) => id
const getCommentsByQuestionID = createSelector(
    [getComments, getQuestionID],
    (comments, questionID) => {
        if (!questionID || !comments || comments.length<=0) return [];
        return comments.filter(comment => comment.post_id == questionID)
    }
)


export default getCommentsByQuestionID;